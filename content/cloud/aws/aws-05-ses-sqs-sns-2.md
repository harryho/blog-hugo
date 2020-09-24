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
       subgraph Notifications
           Bounce_Notification 
           Complaint_Notification 
           Delivery_Notification 
       end
    end

{{</mermaid >}}


