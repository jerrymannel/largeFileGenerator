# largeFileGenerator

> Generate JSON, XML and CSV files

Sample utility to generate large data files.

# Usage

```sh
./largeFileGenerator <opFileName> <type> <size>
Parameters:
	opFileName  :	Output file name
	type        :	json/xml/csv/bin
	size        :	Multiples of 1KB
```

e.g.

To gererate a 1KB CSV file

> ./largeFileGenerator out.csv csv 1

# strings.txt

This file has 3 lines,

*Line 1* : JSON data that will be used for generating the output file.

*Line 2* : XML data that will be used for generating the output file. The first part of the XML data is the name of the root tag

*Line 3* : CSV data that will be used for generating the output file.

~ 🍻~