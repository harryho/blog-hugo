+++
title = "AWS: SQS,SNS,SES - 2"
description = "Use Case - SQS, SNS, SES"
weight=7
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

* Sender code of sample

```java
package email.sample;

// import ....

public class SesSample {

    static final String FROM = "sender@test.com";

    static final String TO = "recipient@test.com";

    static final String CONFIGSET = "ConfigSet";

    // The subject line for the email.
    static final String SUBJECT = "SES test";

    // The email body for recipients with non-HTML email clients.
    static final String TEXTBODY = "This email was sent through Amazon SES "

    public static void main(String[] args) throws IOException {

        try {
            AmazonSimpleEmailService client = 
                AmazonSimpleEmailServiceClientBuilder.standard()
                // Replace the AWS Region
                    .withRegion(Regions.US_WEST_2).build();
            SendEmailRequest request = new SendEmailRequest()
                .withDestination(
                    new Destination().withToAddresses(TO))
                .withMessage(new Message()
                    .withBody(new Body()
                        .withText(new Content()
                            .withCharset("UTF-8").withData(TEXTBODY)))
                    .withSubject(new Content()
                        .withCharset("UTF-8").withData(SUBJECT)))
                .withSource(FROM);
            client.sendEmail(request);
            System.out.println("Email sent!");
            } catch (Exception ex) {
            System.out.println("Error message: " + ex.getMessage());
        }
    }
}

```

* Sample of SQS code

```java
// .....
public class SqsConsumer {
    
      public void receive(Object message) throws Exception {

        if (message instanceof CamelMessage) {
            String body = ((CamelMessage) message).getBodyAs(String.class, camelContext());
            JsonNode envelope = Json.parse(body);
            if (envelope.has("Message")) {
                JsonNode notification = Json.parse(envelope.get("Message").asText());
                String notificationType = notification.get("notificationType").asText();

                log.debug("Processing email notification: " + notificationType);
                switch (notification.get("notificationType").asText()) {
                    case "Received":
                        received.tell(new EmailActorProtocol.EmailReceived(notification), self());
                        break;
                    case "Bounce":
                        response.tell(new EmailActorProtocol.EmailBounced(notification), self());
                        break;
                    case "Delivery":
                        response.tell(new EmailActorProtocol.EmailDelivered(notification), self());
                        break;
                    case "Complaint":
                        response.tell(new EmailActorProtocol.EmailComplaintReceived(notification), self());
                        break;
                    default:
                        throw new RuntimeException(String.format("Notification type %s not supported", notificationType));
                }
            }
        }
    }
}

// EmailActorProtocol 
// public class EmailActorProtocol {
//     public interface EmailResponse {
//     }
//     @Data
//     static public class EmailReceived {
//         private final JsonNode body;
//     }
//     @Data
//     static public class EmailDelivered implements EmailResponse {
//         private final JsonNode body;
//     }
//     @Data
//     static public class EmailBounced implements EmailResponse {
//         private final JsonNode body;
//     }
//     @Data
//     static public class EmailComplaintReceived implements EmailResponse {
//         private final JsonNode body;
//     }
//     static public class GetHealth {
//     }
//     @Data
//     static public class Health {
//         private final boolean healthy;
//     }
// }


```


