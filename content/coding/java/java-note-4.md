+++

title = "Java Note - 4: Date Time "
description="Java 8 provides a comprehensive Date-Time API to work with date, time, and datetime"
+++

## Date-Time API

Through the java.time packages, Java 8 provides a comprehensive Date-Time API to work with date, time, and datetime. By default, most of the classes are based on the ISO-8601 standards. The main classes are

* Instant
   * represents an instant on the timeline and it is suitable for machines, for example, as timestamps for event
* LocalDate,  LocalTime, LocalDateTime
    * represents human readable date, time, and datetime without a time zone. 

* OffsetTime, OffsetDateTime
    * It represent a time and datetime with a zone offset from UTC. 
    
* ZonedDateTime
    * It represents a datetime for a time zone with zone rules, which will adjust the time according to the daylight saving time changes in the time zone.
    
### ISO-8601 Standards for Datetime
* [date]T[time][zone offset] 
* A date component consists of three calendar fields: year, month, and day. Two fields in a date are separated by a hyphen: year-month-day\
* Epoch is Midnight January 1, 1970 UTC

### Useful Datetime-Related Enums
* Month
* DayOfWeek
* ChronoField
* ChronoUnit

## Period 

A period is an amount of time defined in terms of calendar fields years, months, and days. A duration is also an amount of time measured in terms of seconds and nanoseconds. Negative periods are supported. What is the difference between a period and a duration? A duration represents an exact number of nanoseconds, whereas a period represents an inexact amount of time. A period is for humans what a duration is for machines.

## Partial 
Partials
A partial is a date, time, or datetime that does not fully specify an instant on a timeline, but still makes sense to humans. With some more information, a partial may match multiple instants on the timeline.

## Adjusting Dates

Sometimes you want to adjust a date and time to have a particular characteristic, for example, the first Monday of the month, the next Tuesday, etc. You can perform adjustments to a date and time using an instance of the `TemporalAdjuster` interface. The interface has one method, adjustInto(), that takes a Temporal and returns a `Temporal`.

## Formatting

The most important point to keep in mind is that formatting and parsing are
always performed by an object of the DateTimeFormatter class. 

### DateTimeApiDemo

```java
import java.time.Duration;
import java.time.DayOfWeek;

import java.time.format.DateTimeFormatter;
import java.time.format.DateTimeParseException;

import java.time.Instant;
import java.time.LocalDate;
import java.time.LocalTime; 
import java.time.LocalDateTime; 

import java.time.Month;
import java.time.MonthDay;

import java.time.OffsetDateTime;
import java.time.Period;


import java.time.temporal.Temporal;
import java.time.temporal.TemporalAdjusters;
import java.time.temporal.ChronoField;
import java.time.temporal.TemporalAccessor;


import java.time.Year;
import java.time.YearMonth;

import java.time.ZoneOffset;
import java.time.ZonedDateTime;
import java.time.ZoneId;

import java.util.Locale;
import java.util.Set;


import static java.time.Month.JANUARY;
import static java.time.temporal.ChronoUnit.DAYS;
import static java.time.temporal.ChronoUnit.HOURS;
import static java.time.temporal.ChronoUnit.MINUTES;


public class DateTimeApiDemo {

    
    public static String format(Temporal co, String pattern) {
        DateTimeFormatter fmt = DateTimeFormatter.ofPattern(pattern, Locale.US);
        return fmt.format(co);
    }

    public static void parseStr(DateTimeFormatter formatter, String text) {
        try {
            TemporalAccessor ta = formatter.parseBest(text, OffsetDateTime::from, LocalDateTime::from, LocalDate::from);
            if (ta instanceof OffsetDateTime) {
                OffsetDateTime odt = OffsetDateTime.from(ta);
                System.out.println("OffsetDateTime: " + odt);
            } else if (ta instanceof LocalDateTime) {
                LocalDateTime ldt = LocalDateTime.from(ta);
                System.out.println("LocalDateTime: " + ldt);
            } else if (ta instanceof LocalDate) {
                LocalDate ld = LocalDate.from(ta);
                System.out.println("LocalDate: " + ld);
            } else {
                System.out.println("Parsing returned: " + ta);
            }
        } catch (DateTimeParseException e) {
            System.out.println(e.getMessage());
        }
    }

    public static void main(String[] args) {

        // Get current date, time, and datetime 
        LocalDate dateOnly = LocalDate.now();  // 2016-03-12
        LocalTime timeOnly = LocalTime.now();  // 09:17:56.200
        LocalDateTime dateTime = LocalDateTime.now(); // 2016-03-12T09:17:56.200
        ZonedDateTime dateTimeWithZone = ZonedDateTime.now(); // 2016-03-12T09:17:56.202+11:00[Australia/Sydney]

        //  ofXXX() method
        LocalDate ld1 = LocalDate.of(2012, 5, 2); // 2012-05-02
        LocalDate ld2 = LocalDate.of(2012, Month.JULY, 4); // 2012-07-04
        LocalDate ld3 = LocalDate.ofEpochDay(2002); // 1975-06-26
        LocalDate ld4 = LocalDate.ofYearDay(2014, 40); // 2014-02-09

        // The plusXXX( ) and minusXXX( ) Methods
        LocalDate ld = LocalDate.of(2015, 5, 2); // 2015-05-02
        LocalDate ldp1 = ld.plusDays(5); // 2015-05-07
        LocalDate ldp2 = ld.plusMonths(3); // 2015-08-02
        LocalDate ldp3 = ld.plusWeeks(3); // 2015-05-23
        LocalDate ldm1 = ld.minusMonths(7); // 2014-10-02
        LocalDate ldm2 = ld.minusWeeks(3); // 2015-04-11

        // Instant 
        Instant i1 = Instant.ofEpochSecond(20); // i1:1970-01-01T00:00:20Z
        Instant i2 = Instant.ofEpochSecond(55); // i2:1970-01-01T00:00:55Z

        Duration d1 = Duration.ofSeconds(55);
        Duration d2 = Duration.ofSeconds(-17);

        // Compare instants
        System.out.println("i1.isBefore(i2):" + i1.isBefore(i2)); // i1.isBefore(i2):true
        System.out.println("i1.isAfter(i2):" + i1.isAfter(i2)); // i1.isAfter(i2):false

        // Add and subtract durations to instants
        Instant i3 = i1.plus(d1);
        Instant i4 = i2.minus(d2);
        System.out.println("i1.plus(d1):" + i3); // i1.plus(d1):1970-01-01T00:01:15Z

        System.out.println("i2.minus(d2):" + i4);  // i2.minus(d2):1970-01-01T00:01:12Z

        // Add two durations
        System.out.println("d1.plus(d2):" + d1.plus(d2)); // d1.plus(d2):PT38S


        // Print All Zone Id
        Set<String> zoneIds = ZoneId.getAvailableZoneIds();
        for (String zoneId: zoneIds) {
          System.out.println(zoneId);
        }

        // DayOfWeek
        DayOfWeek dw1 = DayOfWeek.from(ld); // THURSDAY 

        // Chrono 

        LocalDateTime now = LocalDateTime.now();
        System.out.println("Year: " + now.get(ChronoField.YEAR));
        System.out.println("Month: " + now.get(ChronoField.MONTH_OF_YEAR));
        System.out.println("Day: " + now.get(ChronoField.DAY_OF_MONTH));
        System.out.println("Hour-of-day: " + now.get(ChronoField.HOUR_OF_DAY));
        System.out.println("Hour-of-AMPM: " + now.get(ChronoField.HOUR_OF_AMPM));
        System.out.println("AMPM-of-day: " + now.get(ChronoField.AMPM_OF_DAY));

        Period p1 = Period.of(2, 3, 5); // 2 years, 3 months, and 5 days
        Period p2 = Period.ofDays(25); // 25 days
        Period p3 = Period.ofMonths(-3); // -3 months
        Period p4 = Period.ofWeeks(3); // 3 weeks (21 days)

        // Date Adjuster
        LocalDate ld3 = ld1.with(TemporalAdjusters.next(DayOfWeek.MONDAY));
        System.out.println(ld3);
        ld3 = ld1.with(TemporalAdjusters.nextOrSame(DayOfWeek.TUESDAY));
        System.out.println(ld3);

        // Date Time Format         
        System.out.println(format(ld, "M/d/yyyy"));
        System.out.println(format(ld, "MM/dd/yyyy"));
        System.out.println(format(ld, "MMM dd, yyyy"));
        System.out.println(format(ld, "MMMM dd, yyyy"));
        System.out.println(format(ld, "EEEE, MMMM dd, yyyy"));
        System.out.println(format(ld, "'Month' q 'in' QQQ"));
        System.out.println(format(ld, "[MM-dd-yyyy][' at' HH:mm:ss]"));

        // Parse date time
        DateTimeFormatter parser = DateTimeFormatter.ofPattern("yyyy-MM-dd['T'HH:mm:ss[Z]]");
        parseStr(parser, "2012-05-31"); //  LocalDate: 2012-05-31
        parseStr(parser, "2012-05-31T16:30:12"); // LocalDateTime: 2012-05-31T16:30:12
        parseStr(parser, "2012-05-31T16:30:12-0500"); // OffsetDateTime: 2012-05-31T16:30:12-05:00
        parseStr(parser, "2012-05-31Hello"); // Text '2012-05-31Hello' could not be parsed, unparsed text found at index 10



    }
}

```
