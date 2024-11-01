IntersectionOfTwoLinkedList: IntersectionOfTwoLinkedList.o
	g++ -g -o IntersectionOfTwoLinkedList.exe IntersectionOfTwoLinkedList.o -pthread    

IntersectionOfTwoLinkedList.o: IntersectionOfTwoLinkedList/IntersectionOfTwoLinkedList.cpp
	g++ -g  -c -pthread IntersectionOfTwoLinkedList/IntersectionOfTwoLinkedList.cpp
