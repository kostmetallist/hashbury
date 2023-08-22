#define _CRTDBG_MAP_ALLOC
#include <crtdbg.h>
#include <iostream>
#include "entry.h"
#include "hashmap.h"

int main()
{
    Entry *e1 = new Entry(1, 42);
    std::cout << "The interaction with the data structure is expected here.\n";
    delete e1;

    Hashmap *hm = new Hashmap(8);
    delete hm;

    _CrtSetReportMode(_CRT_WARN, _CRTDBG_MODE_DEBUG);
    _CrtDumpMemoryLeaks();
    return 0;
}
