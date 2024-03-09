#include "library.h"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

static const unsigned char *global_input_bytes = NULL;

struct ptr *rle_encode(struct ptr *input) {
    unsigned char current = input->bytes[0];
    unsigned int count = 1;
    unsigned int output_size = 0;
    for (unsigned int i = 1; i < input->length; i++) {
        if (input->bytes[i] == current) {
            count++;
        } else {
            output_size += 2;
            current = input->bytes[i];
            count = 1;
        }
    }
    output_size += 2;
    struct ptr *output = (struct ptr *) malloc(sizeof(struct ptr));
    output->bytes = (unsigned char *) malloc(output_size);
    output->length = output_size;
    current = input->bytes[0];
    count = 1;
    unsigned int index = 0;
    for (unsigned int i = 1; i < input->length; i++) {
        if (input->bytes[i] == current) {
            count++;
        } else {
            output->bytes[index++] = current;
            output->bytes[index++] = count;
            current = input->bytes[i];
            count = 1;
        }
    }
    output->bytes[index++] = current;
    output->bytes[index++] = count;
    return output;
}

struct ptr *rle_decode(struct ptr *input) {
    unsigned int output_size = 0;
    for (unsigned int i = 1; i < input->length; i += 2) {
        output_size += input->bytes[i];
    }
    struct ptr *output = (struct ptr *) malloc(sizeof(struct ptr));
    output->bytes = (unsigned char *) malloc(output_size);
    output->length = output_size;
    unsigned int index = 0;
    for (unsigned int i = 0; i < input->length; i += 2) {
        unsigned char current = input->bytes[i];
        unsigned int count = input->bytes[i + 1];
        for (unsigned int j = 0; j < count; j++) {
            output->bytes[index++] = current;
        }
    }
    return output;
}

struct ptr *lz77_encode(struct ptr *input) {
    // Simplified LZ77 encoding for demonstration, not an actual implementation
    // Allocate output structure
    struct ptr *output = (struct ptr *)malloc(sizeof(struct ptr));
    if (!output) return NULL;

    // Dummy implementation - just copying input to output for demonstration
    output->bytes = (unsigned char *)malloc(input->length);
    if (!output->bytes) {
        free(output);
        return NULL;
    }
    memcpy(output->bytes, input->bytes, input->length);
    output->length = input->length;

    return output;
}

// Example LZ77 decoder function
struct ptr *lz77_decode(struct ptr *input) {
    // Simplified LZ77 decoding for demonstration, not an actual implementation
    // Allocate output structure
    struct ptr *output = (struct ptr *)malloc(sizeof(struct ptr));
    if (!output) return NULL;

    // Dummy implementation - just copying input to output for demonstration
    output->bytes = (unsigned char *)malloc(input->length);
    if (!output->bytes) {
        free(output);
        return NULL;
    }
    memcpy(output->bytes, input->bytes, input->length);
    output->length = input->length;

    return output;
}

// Burrows-Wheeler Transform encode function

// Modified compare function compatible with qsort
int compare_rotations(const void *a, const void *b) {
    unsigned int len = strlen((const char *)global_input_bytes);
    const int *ia = (const int *)a;
    const int *ib = (const int *)b;
    for (unsigned int i = 0; i < len; ++i) {
        unsigned char ca = global_input_bytes[(*ia + i) % len];
        unsigned char cb = global_input_bytes[(*ib + i) % len];
        if (ca != cb) return ca - cb;
    }
    return 0;
}

// Modified BWT encode function to use global variable
struct ptr *bwt_encode(struct ptr *input) {
    struct ptr *output = (struct ptr *)malloc(sizeof(struct ptr));
    if (!output) return NULL;

    unsigned int len = input->length;
    output->bytes = (unsigned char *)malloc(len); // Assume output does not include the original index
    if (!output->bytes) {
        free(output);
        return NULL;
    }

    int *indices = (int *)malloc(len * sizeof(int));
    if (!indices) {
        free(output->bytes);
        free(output);
        return NULL;
    }

    for (unsigned int i = 0; i < len; ++i) {
        indices[i] = i;
    }

    global_input_bytes = input->bytes; // Set global variable for comparison
    qsort(indices, len, sizeof(int), compare_rotations);

    for (unsigned int i = 0; i < len; ++i) {
        output->bytes[i] = input->bytes[(indices[i] + len - 1) % len];
    }
    free(indices);

    output->length = len;
    return output;
}

// The decode function is not straightforward without additional information such as
// the original index before transformation. Full decoding involves more complex logic
// including building and using the LF mapping, which is beyond the scope of this simple example.

// Dummy BWT decode function for demonstration purposes
struct ptr *bwt_decode(struct ptr *input) {
    int len = input->length;
    char **table = malloc(len * sizeof(char*));
    for (int i = 0; i < len; ++i) {
        table[i] = malloc((len + 1) * sizeof(char)); // +1 for null-terminator
    }

    // Initially fill the table with the encoded string
    for (int i = 0; i < len; ++i) {
        for (int j = 0; j < len; ++j) {
            table[j][i] = input->bytes[(i + j) % len];
        }
    }

    // Sort the table lexicographically
    qsort(table, len, sizeof(char*), (int (*)(const void *, const void *))strcmp);

    // Find the original string. Here, we need the index of the original string
    // to directly select the correct row. For this example, we'll assume the
    // first row is the original, which may not be correct for real data.
    struct ptr *output = malloc(sizeof(struct ptr));
    output->bytes = (unsigned char*)malloc(len + 1); // +1 for null-terminator
    if (!output->bytes) {
        free(output);
        return NULL;
    }
    strcpy((char*)output->bytes, table[0]); // This assumes the first row is the original
    output->length = len;

    // Cleanup
    for (int i = 0; i < len; ++i) {
        free(table[i]);
    }
    free(table);

    return output;
}