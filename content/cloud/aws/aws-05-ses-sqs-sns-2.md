+++
title = "AWS: SQS,SNS,SES - 2"
description = "Use Case - SQS, SNS, SES"
weight=7
draft=true
+++

## Use Case

### Overview

{{<mermaid>}}
graph LR
    Sender_Email("test@test.com")
    Email_Failed
    Email_Delivered
    SNS_Subscriptions --> Email_Failed
    SNS_Subscriptions --> Email_Delivered
    Bounce_Notification --> Email_Failed
    Complaint_Notification --> Email_Failed
    Delivery_Notification --> Email_Delivered
    subgraph SQS     
        subgraph Email_Status_Queue
            SNS_Subscriptions
        end
    end 
    subgraph SNS
       subgraph Topics
           Email_Failed
           Email_Delivered
       end
    end
    subgraph SES
       Sender_Email
       subgraph Notifications
           Bounce_Notification 
           Complaint_Notification 
           Delivery_Notification 
       end
    end

{{</mermaid >}}


### SNS Setup

* Create a topic for failed email, e.g. bounce or spam complaint
    - It is named **Email_Failed** in the diagram above

* Create a topic for delivered email
    - It is named **Email_Delivered** in the diagram above


### SES Setup

* Create a new domain for sender email, e.g. test@test.com
  
* Setup the notifications
    - Bounce Notification maps to **Email_Failed**
    - Complaint Notification maps to **Email_Failed**
    - Delivery Notification maps to **Email_Delivered**
  
* Verify the domain - test.com

* Verify the DKIM - *.domainkey.test.com

### SQS Setup

* Create a new queue named **Email_Status_Queue**
* Add SNS subscriptions **Email_Failed** and **Email_Delivered** to the queue
  

### Integration


```java
package email.sample;

// import ....

public class AmazonSESSample {

  static final String FROM = "sender@example.com";

  static final String TO = "recipient@example.com";


  static final String CONFIGSET = "ConfigSet";

  // The subject line for the email.
  static final String SUBJECT = "Amazon SES test (AWS SDK for Java)";
  
  // The HTML body for the email.
  static final String HTMLBODY = "<h1>Amazon SES test (AWS SDK for Java)</h1>"
      + "<p>This email was sent with <a href='https://aws.amazon.com/ses/'>"
      + "Amazon SES</a> using the <a href='https://aws.amazon.com/sdk-for-java/'>" 
      + "AWS SDK for Java</a>";

  // The email body for recipients with non-HTML email clients.
  static final String TEXTBODY = "This email was sent through Amazon SES "
      + "using the AWS SDK for Java.";

  public static void main(String[] args) throws IOException {

    try {
      AmazonSimpleEmailService client = 
          AmazonSimpleEmailServiceClientBuilder.standard()
          // Replace US_WEST_2 with the AWS Region you're using for
          // Amazon SES.
            .withRegion(Regions.US_WEST_2).build();
      SendEmailRequest request = new SendEmailRequest()
          .withDestination(
              new Destination().withToAddresses(TO))
          .withMessage(new Message()
              .withBody(new Body()
                  .withHtml(new Content()
                      .withCharset("UTF-8").withData(HTMLBODY))
                  .withText(new Content()
                      .withCharset("UTF-8").withData(TEXTBODY)))
              .withSubject(new Content()
                  .withCharset("UTF-8").withData(SUBJECT)))
          .withSource(FROM)
          // Comment or remove the next line if you are not using a
          // configuration set
          .withConfigurationSetName(CONFIGSET);
      client.sendEmail(request);
      System.out.println("Email sent!");
    } catch (Exception ex) {
      System.out.println("The email was not sent. Error message: " 
          + ex.getMessage());
    }
  }
}

```


