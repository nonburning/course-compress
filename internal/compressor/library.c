#include "library.h"

#include <stdio.h>
#include <stdlib.h>

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
    return input;
}

struct ptr *lz77_decode(struct ptr *input) {
    return input;
}

struct ptr *bwt_encode(struct ptr *input) {
    return input;
}

struct ptr *bwt_decode(struct ptr *input) {
    return input;
}
