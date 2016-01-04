# topd
make simple search querys against a TOPdesk DB in your console

### Example:

Search for a device by it's hostname:

Usage:  topd <search>

    Examples:
        topd nb2737         //search by inventory number

Example:

`topd NB113`

output:

    +----------+--------------------+----------------+----------------------------+
    | HOSTNAME |        USER        |      TYPE      |       SPECIFICATION        |
    +----------+--------------------+----------------+----------------------------+
    | NB1131   | IT                 | Latitude E6430 | Core i5-3230M, 8GB RAM,... |
    | NB1132   | Franz              | Latitude E6430 | Core i5-3230M, 8GB RAM,... |
    | NB1133   | Allgemein          | Latitude E6430 | Core i5-3230M, 8GB RAM,... |
    | NB1134   | Markus             | Latitude E6430 | Core i5-3230M, 8GB RAM,... |
    | NB1136   | Beat               | Latitude E7440 | Core i5-4300U, 8GB RAM,... |
    | NB1137   | Frederik           | Latitude E7440 | Core i5-4300U, 8GB RAM,... |
    | NB1139   | Mustafa            | Latitude E7440 | Core i5-4300U, 8GB RAM,... |
    +----------+--------------------+----------------+----------------------------+