# Ref Store

A typical deduplicating storage stores data object only once. In place of each occurrence of that data object, only a reference to the actual data object is stored. 

Let’s take an example of a block file system. Assume fixed sized blocks of 1 MB size. Let’s consider a file where a same block occurs 2 times, once at offset 3MB and once at offset 5MB. File system can store the data object only once. And filesystem metadata will contain the pointers to the actual data blocks at each offset 3MB and 5MB. These pointers are called forward references. 

Such deduplicating storage system also need back references as well. These back references are used
- as a replacement of reference counts - i.e. when all the back references are deleted, the data object becomes eligible for deletion.
- by file system check utility to make file system consistent (or cleanup unwanted data objects) after a crash. 

Forward references can be as simple as unique integer values - typically generated using counter variables. As long as the file system guarantees uniqueness of data object to counter value mapping, the consistency can be guaranteed. 

Modern performance efficient applications (especially database applications) can make use of a library that provides above mentioned deduplication functionality.
