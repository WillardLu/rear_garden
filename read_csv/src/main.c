/// @copyright Copyright 2023 Willard Lu
///
/// Use of this source code is governed by an MIT-style
/// license that can be found in the LICENSE file or at
/// https://opensource.org/licenses/MIT.

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "tools.h"

int main(int argc, char *argv[]) {
  char *filename = argv[1];
  // 打开CSV文件
  FILE *fp = fopen(filename, "r");
  if (fp == NULL) {
    printf("open file failed\n");
    return -1;
  }
  // 获取文件内容大小
  fseek(fp, 0L, SEEK_END);
  long fsize = ftell(fp);
  fseek(fp, 0L, SEEK_SET);
  // 分配空间
  char *buf = (char *)calloc(fsize + 1, sizeof(char));
  // 找到第一个逗号所在位置
  for (int i = 0; i < fsize; i++) {
    if (fgetc(fp) == ',') {
      break;
    }
  }
  // 读取第一个回车符之前的所有字符，最后加上一个逗号
  int regex_size = 0; // 记录实际内容长度
  char c1 = 0;
  for (int i = 0; i < fsize; i++) {
    c1 = fgetc(fp);
    regex_size++;
    if (c1 == '\n') {
      buf[i] = ',';
      break;
    }
    buf[i] = c1;
  }
  printf("%s\n", buf);
  printf("size = %d\n", regex_size);
  // 关闭文件
  fclose(fp);
  fp = NULL;
  //
  char *regex_temp1;
  char *regex_temp2;
  int quote_num = 0;
  regex_temp1 = (char *)calloc(regex_size + 1, sizeof(char));
  regex_temp2 = (char *)calloc(regex_size + 1, sizeof(char));
  for (int i = 0; i < regex_size; i++) {
    c1 = buf[i];
    regex_temp1[i] = c1;
    if (c1 == '"') {
      quote_num++;
      continue;
    }
    if (c1 != ',') {
      continue;
    }
    if (quote_num > 0) {
      if (quote_num % 2 != 0) {
        continue;
      }
      // 去掉字符串两端的双引号和最后的逗号
      for (int j = 1; j < i - 1; j++) {
        regex_temp2[j - 1] = regex_temp1[j];
      }
      // 把字符串中连续的两个双引号替换成一个双引号
      for (int j = 0; j < i - 1; j++) {
      }
    }
    printf("regex_temp2 = %s\n", regex_temp2);
    quote_num = 0;
  }
  // 释放给buf分配的内存空间
  free(buf);
  free(regex_temp1);
  free(regex_temp2);

  return 0;
}
