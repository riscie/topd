# topd
[![Go Report Card](http://goreportcard.com/badge/riscie/topd)](http://goreportcard.com/report/riscie/topd)


**make simple search querys against a TOPdesk DB in your console**

#### Usage:
    topd <search keyword>


#### Example Output:

![sample output png](http://langhard.com/github/topd.png "sample topd output")


#### Search Examples:

**search by TOPdesk inventory number or hostname**

    topd PC2115

**search by IP Address (finds all devices with an IP startig with 10.10.3.)**

    topd 10.10.3.

**search by MAC address**

    topd 01-00-5e-7f-ff-fa
    
**serach by the users name**

    topd Max

**(Shows only active devices)**

