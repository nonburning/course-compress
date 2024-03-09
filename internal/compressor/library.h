#ifndef TPRO_LIBRARY_LIBRARY_H
#define TPRO_LIBRARY_LIBRARY_H

struct ptr{
    unsigned char *bytes;
    unsigned int length;
};

int compare_rotations(const void *a, const void *b);

struct ptr *rle_encode(struct ptr *input);

struct ptr *rle_decode(struct ptr *input);

struct ptr *lz77_encode(struct ptr *input);

struct ptr *lz77_decode(struct ptr *input);

struct ptr *bwt_encode(struct ptr *input);

struct ptr *bwt_decode(struct ptr *input);

#endif //TPRO_LIBRARY_LIBRARY_H
