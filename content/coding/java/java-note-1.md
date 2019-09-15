+++

title = "Java Note - 1: Enum"
description="Replace constant property of the interface or abstract class with Enum"
+++


## Prerequisites

>*Java 1.5+*

## New type: Enum

Enum was a great improvement in Java 1.5. From that more and more developer abandom the interface or abstract class as constant variable container. 

### Before Java 5

Before Java 1.5 you will following coding in many Java program. 

```java
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

Above program looks very good. Please take a close look and check it carefully. You will find the program will never return __*en-au*__, since there is a typo in the constant AU. It should be __*Australia*__ instead of __*Australian*__. I believe many developers have short sight problem like me, and it happened again and again. Using string as constant flag is not a good option, but there is no other better solution before Java 1.5.

### After Java 5

After Java 1.5, you will see the change below. Enum is the best container for constants. It can help you check the program time. Meanwhile, it can simplfy your coding. 



```java
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

After Java Java 1.7, there are some new features. One of these new features is Switch statement. Now it supports String. It is a great for Java developer.With this new feature, the old Enum can be enhanced and the Util class can provide more handy methods (Overload method getLanguageCode) for development. 

New Enum class can support flexible requirement. In the early version of Enum, the toString method only will return exactly the specified constanct name. Now it can be overrided with toString to return different constant name. 

```java
// It can return customized name and simplify coding 
public enum Country {
	AU("Australia", "au","en-au"), 
	UK("United Kingdom", "en-uk"),
	US("United State", "us","en-us");

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

```java
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
