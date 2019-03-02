/**
 * morpheme data structure
 */
typedef struct khaiii_morph_t_ {
    const char* lex;    ///< lexical
    const char* tag;    ///< part-of-speech tag
    int begin;    ///< morpheme begin position
    int length;    ///< morpheme length
    char reserved[8];    ///< reserved
    const struct khaiii_morph_t_* next;    ///< next pointer
} khaiii_morph_t;
 
/**
 * word data structure
 */
typedef struct khaiii_word_t_ {
    int begin;    ///< word begin position
    int length;    ///< word length
    char reserved[8];    ///< reserved
    const khaiii_morph_t* morphs;    ///< morpheme list
    const struct khaiii_word_t_* next;    ///< next pointer
} khaiii_word_t;

char* init_funct_ptrs(char* libpath);
 
char* khaiii_version();

int khaiii_open(const char* rsc_dir, const char* opt_str);

const khaiii_word_t* khaiii_analyze(int handle, const char* input, const char* opt_str);

void khaiii_free_results(int handle, const khaiii_word_t* results);

void khaiii_close(int handle);

const char* khaiii_last_error(int handle);