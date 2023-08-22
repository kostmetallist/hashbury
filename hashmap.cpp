#include "hashmap.h"
#include "entry.h"

void Hashmap::initialize_table() {
	table = new Entry *[table_size];
	for (int i = 0; i < table_size; ++i) {
		table[i] = nullptr;
	}
}

Hashmap::Hashmap() {
	initialize_table();
}

Hashmap::Hashmap(const int table_size) {
	this->table_size = table_size;
	initialize_table();
}

Hashmap::~Hashmap() {
	for (int i = 0; i < table_size; ++i) {
		if (table[i] != nullptr) {
			delete table[i];
		}
	}

	delete[] table;
}
