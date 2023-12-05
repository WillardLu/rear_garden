// @copyright Copyright 2023 Willard Lu
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

#include <stdio.h>
#include <string.h>
#include <pcre.h>

int main() {
  pcre *re; // 保存编译后的正则表达式
  const char *error; // 保存错误信息
  int erroffset; // 保存错误位置
  int rc; // 保存匹配串的偏移位置
  const char *pattern = "-.+[a]+."; // 示例用正则表达式
  // 1. 编译正则表达式
  re = pcre_compile(pattern, 0, &error, &erroffset, NULL);
  if (re == NULL) {
    // 输出错误信息
    fprintf(stderr, "PCRE compilation failed at offset %d: %s\n", erroffset, error);
    return 1;
  }
  // 2. 模式匹配
  const char *subject = "--===abcfooabcfoo";
  rc = pcre_exec(re, NULL, subject, strlen(subject), 0, 0, NULL, 0);
  if (rc < 0) {
    if (rc == PCRE_ERROR_NOMATCH) {
      printf("No match\n");
    }
    else {
      printf("PCRE execution failed at offset %d: %d\n", erroffset, rc);
    }
    pcre_free(re);
    return 1;
  }
  printf("Matched %d groups\n", rc);
  // 释放内存
  pcre_free(re);
  return 0;
}