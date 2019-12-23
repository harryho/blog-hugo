+++

title = "Java Note - 3: Path and Files"
description="Some good part of Java 7 - Path and Files API"
+++


## Prerequisites

>*Java 7+*

## Good stuff from not shiny Java 7

If you are planning to refactor your code, please give a second thought. It is time to dump to try these new features. When the Java 7 was released, I was kind of disappointed without lambda, jigsaw as most developers, but when I tried new Path, Files API, I found that is great improvement. The enhancement of this new IO is really useful. It save so much effort for Java developer. 

To be hoenst, before Java 7, Coding file manipulation in Java is very headache task. I say "headache" it doesn't mean it is difficult. Just comparing with other program lanugage, you had to take much more effort to take care of the boilerplate, and all are tedious job. That is why sometimes I prefer cmd in Window or bash in Linux to complete the task instead of using Java to handle file manipulation. Now I think I can refactor old file manipulation coding and make it much more elegant. 


### Better file visitor implementation

Following is simple customizaed file visitor which has been the part of my old util.

```java
import java.io.IOException;
import java.nio.file.FileVisitResult;
import java.nio.file.Path;
import java.nio.file.SimpleFileVisitor;
import java.nio.file.attribute.BasicFileAttributes;

public class CustomFileVisitor extends SimpleFileVisitor<Path> {
    @Override
    public FileVisitResult postVisitDirectory(Path dir , IOException arg1) throws IOException {
        System.out.println( "post visit dir :  "+  dir );
        return FileVisitResult.CONTINUE;
        }

        @Override
    public FileVisitResult preVisitDirectory(Path dir , IOException arg1) throws IOException {
        System.out.println( "post visit dir :  "+  dir );
        return FileVisitResult.CONTINUE;
        }

    @Override
    public FileVisitResult visitFile(Path file, BasicFileAttributes attr)
            throws IOException {
         if ( attr.isSymbolicLink() )
          { System.out.println( " symbolic link  :  "+  file  );
          
          }else if (  attr.isSymbolicLink() ){              
              System.out.println( " regular file : "+ file );
          }
        return FileVisitResult.CONTINUE;
    }
     
    @Override
    public FileVisitResult visitFileFailed(Path file, IOException exc)
            throws IOException {        
         System.err.println( exc.getMessage());
        return FileVisitResult.CONTINUE;
    }
}
```

To use this customized  is so easy. Just 3 lines coding you can test it by yourself. 
```
        CustomFileVisitor fileVisitor = new CustomFileVisitor();
        Path path = Paths.get("TestDir");
        Files.walkFileTree(path, fileVisitor);
```

### ARM 

Automatic resource management is another attractive features of Java 7 and project coin. As name itself implies that now JVM is going to be handling all the external resource and make programmer free to bother about resource management, especially for people like me miss the `using` statement in C#. Sometimes I wonder why Java is such stubborn not to learn some good features from C#. As we know, C# comes after Java and copies most concept at the early stage, but it really pushed Object Oriented Concept (OOC) to a new level and inspired Java world a lot with its many good feature. I really hope someday I can code in Java as simple as C#. Wise men learn by other men's mistakes; fools by their own. 

In the past, java programmers use any external resources like file, printer or any devices to close after my program execution complete. Normally we close the resources which we have open in beginning of our program or we decide that if program finish normally how to manage the resource or if our program finish abnormally how to close the resource. Following are comparison of old and new style. 

*Snippet of old style*

```java
FileInputStream exchangeCurrencyReader= null;
FileOutputStream exchangeCurrencyWriter = null;
try {
    exchangeCurrencyReader = new FileInputStream("AUDvsUSD.txt");
    exchangeCurrencyWriter = new FileOutputStream("AUDvsUSD.txt");
    int var;
    while (var = exchangeCurrencyReader.read()) != -1)
        exchangeCurrencyWriter.write(var);
} 
finally {
    if (exchangeCurrencyReader!= null)
        exchangeCurrencyReader.close();
    if (exchangeCurrencyWriter!= null)
        exchangeCurrencyWriter.close();
}
```

*Code in Java 7*

```java
try ( FileInputStream exchangeCurrencyReader = new FileInputStream("AUDvsUSD.txt");
    FileOutputStream exchangeCurrencyWriter = new FileOutputStream("AUDvsUSD.txt")){
      int var;
      while((var= exchangeCurrencyReader.read()) != -1 )
            exchangeCurrencyWriter.write();
}

```

In the code above we have declare two file stream one is input file we are reading from one file and writing to another file. After the whole process both streams will be closed automatically either the code has been executed properly or not. Both exchangeCurrencyReader.close() and exchangeCurrencyWriter.close() methods will be called automatically which is the best part of ARM. We should not miss good part from Java 7. 


### New file change monitor service

After some homework for new features of Java 7, I am tring to use file watch serviice from Java 7 to replace old file monitor program. It is great and quite simple to use. I have updated to production. 

Usually most Java based system somehow needs such monitor function, there will be separate process or thread to run this service, and there should be a call back handler triggered by this service. Everytime there is any file amended, the service will trigger the call back handler to complete some tasks. 

Following is sample of file watching service. I removed call back part which is relevant to the business. 

```java
import java.nio.file.FileSystems;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardWatchEventKinds;
import java.nio.file.WatchEvent;
import java.nio.file.WatchKey;
import java.nio.file.WatchService;
import java.util.HashMap;
import java.util.Map;

public class FileWatchService {

    public static void watchFileUpdate() {
        try (WatchService service = FileSystems.getDefault().newWatchService()) {
            Map<WatchKey, Path> eventMap = new HashMap<>();
            Path dir = Paths.get("TestDir");

            eventMap.put(dir.register(service, StandardWatchEventKinds.ENTRY_MODIFY),dir);

            WatchKey key;
            do {
                key = service.take();
                Path eventPath = eventMap.get(key);
                for (WatchEvent<?> event : key.pollEvents()) {
                    WatchEvent.Kind<?> kind = event.kind();
                    Path path = (Path) event.context();
                    System.out.println(eventPath + " : " + kind + "  : " + path);
                }
            } while (key.reset());
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    
    public static void main (String [] args ){
       watchFileUpdate();
    }
}
```

The sample above shows one kind of events. Actually if you check the API doc, there are another two kinds of event. One is  StandardWatchEventKinds.ENTRY\_CREATE , the other is  StandardWatchEventKinds.ENTRY\_DELETE. These events cover almostly all business requirements.   




































































