+++
tags = ["java"]
categories = ["blog"]
date = "2011-08-10T14:59:31+11:00"
title = "Better practice in Java - Part 1"
draft = true
+++


Enum was a great improvement in Java 1.5. From that more and more developer abandom the interface or abstract class as constant variable container. 

Before Java 1.5 you will following coding in many Java program. 
```
    // Use interface or abstract class as constant variable container
    public interface Country {
         public static final String AU = "Australian";
         public static final String UK = "United Kingdom"; 
         public static final String US = "United State"; 
    }
    
    public class Util {
         
    	public static String getLanguageCode(String country) {
    
    		String languageCode = "en";
    		switch (country) {
    			case Country.AU:
    				languageCode = "en-au";
    				break;
    			case Country.UK:
    				languageCode = "en-uk";
    				break;
    			case Country.US:
    				languageCode = "en-us";
    				break;
		    }
		    return languageCode;
    	}
    }
```


Above program looks very good. Please take a close look and check it carefully. You will find the program will never return <u>"en-au"</u>, since there is a typo in the constant AU. It should be <u>"Australia"</u> instead of <u>"Australian"</u>. I believe many developers have short sight problem like me, and it happened again and again. Using string as constant flag is not a good option, but there is no other better solution before Java 1.5.

After Java 1.5, you will see the change below. Enum is the best container for constants. It can help you check the program time. Meanwhile, it can simplfy your coding. 

```
    // Use Enum as constant variable container
    public enum  Country {
        Australia, UnitedKingdom, UnitedState
    }
    
    public class Util {
         
    	public static String getLanguageCode(Country country) {
    
    		String languageCode = "en";
    		switch (country) {
    			case Australia:
    				languageCode = "en-au";
    				break;
    			case UnitedKingdom:
    				languageCode = "en-uk";
    				break;
    			case UnitedState:
    				languageCode = "en-us";
    				break;
		    }
		    return languageCode;
    	}
    }
```

Now you program will not be by any unintentional typo, since it will throw you compile error before you run the application. If you haven't refactor your static constants container, it is time to improve it now. 

(Above were written in 2006. After Java 1.7, I found it can be improved with new feature of Java 1.7 )

After Java Java 1.7, there are some new features. One of these new features is Switch statement. Now it supports String. It is a great for Java developer.

With this new feature, the old Enum can be enhanced and the Util class can provide more handy methods (Overload method getLanguageCode) for development.

New Enum class, which can support flexible requirement. In the early version of Enum, the toString method only will return exactly the specified constanct name. Now the Enum can be override with toString to return different constant name. 

```
    // It can return customized name and simplify coding 
    public enum Country {
	    AU("Australia", "au","en-au"), UK("United Kingdom", "en-uk"), US("United State", "us","en-us");

    	String countryName;
    	String countryCode;
        String languageCode;
        
    	private Country(String name, String code) {
    		countryName = name;
    		countryCode = code;
    	}
    
    	public String getCode() {
    		return countryCode;
    	}
    	
        public String getLanguageCode() {
    		return languageCode;
    	}
    	
    	@Override
    	public String toString() {
    		return countryName;
    	}
}
```

The Uitl class can convert any country name or country code to Enum Country, vice versa. Now developer can seamless convert the String from UI to the Enum, since on the UI, usually the country name will Australia or United Kingdom instead of just AU or UK. For coding, use AU or UK can simplify coding and is friendly to developer. 
```

public class Util {
        
    public static Country convertCountryNameOrCode(String nameOrCode ) {
	    
		Country country = null;
		switch (nameOrCode) {
			case "au":
			case "AU":
			case "Australia":
				country = Country.AU;
				break;
			case "uk":
			case "UK":	
			case "United Kingdom":
				country = Country.UK;
				break;
			case "us":
			case "US":	
			case "United State":
				country = Country.UK;
				break;
		}
		return country;
	}
	
	public static String getCountryName(  Country country ){
		String countryName = null;
		switch (country){
			case AU:
				countryName = Country.AU.toString();
				break;
			case US:
				countryName = Country.UK.toString();
				break;
			case UK:
				countryName = Country.US.toString();
				break;
			default:
				System.err.println("Unknow Country");
				assert false;
				break;
		}
		return countryName;
	}

	
	public static String getCountryCode(  Country country ){
		String countryCode = null;
		switch (country){
			case AU:
				countryCode = Country.AU.getCode();
				break;
			case US:
				countryCode = Country.UK.getCode();
				break;
			case UK:
				countryCode = Country.US.getCode();
				break;
			default:
				System.err.println("Unknow Country");
				assert false;
				break;
		}
		return countryCode;
	}
        
    public static String getLanguageCode(Country country) {
    
    		String languageCode = "en";
    		switch (country) {
    			case AU:
    				languageCode = Country.AU.getLanguageCode() ;
    				break;
    			case UK:
    				languageCode =  Country.UK.getLanguageCode();
    				break;
    			case US:
    				languageCode =  Country.US.getLanguageCode();
    				break;
		    }
		    return languageCode;
    }
}
```

Are you planning to refactor your code? Just do it. You will love these new features. 



When the Java 7 was released, I was kind of disappointed without lamda, jigsaw as most developers, but when I tried new Path, Files API, I found that is great improvement. The enhancement of this new IO is really useful. It save so much effort for Java developer. 

To be hoenst, before Java 1.7, Coding file manipulation in Java is very headache task. I say "headache" it doesn't mean it is difficult. Just comparing with other program lanugage, you need to take much more effort to do the same job. If there is no limitation, I prefer cmd in Window or bash in Linux to complete the task instead of Java to handle file manipulation. Now I think I can refactor old file manipulation coding and make it much more elegant. 

Coming with the new IO, try-with-resource is another great improvement and I don't need to miss the using statement in C#. Sometimes I wonder why Java is such reluctant to learn some good features from C#. As we all know, C# comes after Java and copies most concept at the early stage, but it really pushed Java to new level and inspired Java world a lot with its many good feature. I really hope someday I can write Java getter and setter in C# coding style. Wise men learn by other men's mistakes; fools by their own.

Following is simple customization which has been the part of my old util.

```
import java.io.IOException;
import java.nio.file.FileVisitResult;
import java.nio.file.Path;
import java.nio.file.SimpleFileVisitor;
import java.nio.file.attribute.BasicFileAttributes;

public class CustomFileVisitor extends SimpleFileVisitor&lt;Path&gt; {
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

To use this customized is so easy. Just 3 lines coding you can test it by yourself. 
```
		CustomFileVisitor fileVisitor = new CustomFileVisitor();
		Path path = Paths.get("TestDir");
		Files.walkFileTree(path, fileVisitor);
```

Next time, I will try watch service. I am planning to use it to replace old file monitor program. 





As last time I mentioned, I am tring to use file watch serviice from Java 1.7 to replace old file monitor program. It is great and quite simple to use. I have updated to production. 

I believe if you system needs such monitor function, there will be separate thread to run this service, and there should be a call back handler triggered by this service. Everytime there is any file amended, the service will trigger the call back handler to complete some tasks. 

Following is sample of file watching service. I removed call back part for our business. 

```
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
			Map&lt;WatchKey, Path&gt; eventMap = new HashMap&lt;&gt;();
			Path dir = Paths.get("TestDir");

			eventMap.put(dir.register(service, StandardWatchEventKinds.ENTRY_MODIFY),dir);

			WatchKey key;
			do {
				key = service.take();
				Path eventPath = eventMap.get(key);
				for (WatchEvent&lt;?&gt; event : key.pollEvents()) {
					WatchEvent.Kind&lt;?&gt; kind = event.kind();
					Path path = (Path) event.context();
					System.out.println(eventPath + " : " + kind + "  : " + path);
				}
			} while (key.reset());
		} catch (Exception e) {

		}
	}
	
	public static void main (String [] args ){
	   watchFileUpdate();
	}
}
```

The sample above shows one kind of events. Actually if you check the API doc, there are another two kinds of event. One is  StandardWatchEventKinds.ENTRY_CREATE , the other is  StandardWatchEventKinds.ENTRY_DELETE. 

What are you waiting for? Refactor your obsolete program to new handy API now. 




































































