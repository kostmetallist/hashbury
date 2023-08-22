#pragma once
#include "entry.h"

const int DEFAULT_SIZE = 2 << 8;

class Hashmap {
	int table_size = DEFAULT_SIZE;
	Entry **table;

	void initialize_table();

public:
	Hashmap();
	Hashmap(const int table_size);
	~Hashmap();
};
