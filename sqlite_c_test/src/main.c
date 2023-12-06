// @copyright Copyright 2023 Willard Lu
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

#include <sqlite3.h>
#include <stdio.h>

int Callback(void *, int, char **, char **);

int main() {
  sqlite3 *db;
  char *err_msg = 0;

  // 1. 打开数据库
  int rc = sqlite3_open("test1.db", &db);
  if (rc != SQLITE_OK) {
    fprintf(stderr, "Can't open database: %s\n", sqlite3_errmsg(db));
    sqlite3_close(db);
    return 1;
  }

  // 2. 执行sql命令
  char *sql = "select * from t1";

  // 使用sqlite3_exec()函数来执行sql命令
  rc = sqlite3_exec(db, sql, Callback, 0, &err_msg);
  if (rc != SQLITE_OK) {
    fprintf(stderr, "Failed to select data\n");
    fprintf(stderr, "SQL error: %s\n", err_msg);
    sqlite3_free(err_msg);
    sqlite3_close(db);
    return 1;
  }

  // 使用sqlite3_prepare_v2预编译命令来执行sql命令
  sqlite3_stmt *res;
  rc = sqlite3_prepare_v2(db, sql, -1, &res, 0);
  if (rc == SQLITE_OK) {
    sqlite3_bind_int(res, 1, 3);
  } else {
    fprintf(stderr, "Failed to prepare statement: %s\n", sqlite3_errmsg(db));
  }
  int step = 0;
  while (1) {
    step = sqlite3_step(res); // 查询一步
    if (step == SQLITE_ROW) {
      printf("%s: ", sqlite3_column_text(res, 0));
      printf("%s\n", sqlite3_column_text(res, 1));
    } else {
      break;
    }
  }

  sqlite3_finalize(res); // 销毁预编译对象
  sqlite3_close(db); // 关闭数据库
  return 0;
}

// 回调函数
int Callback(void *NotUsed, int argc, char **argv, char **azColName) {
  NotUsed = 0;
  for (int i = 0; i < argc; i++) {
    printf("%s = %s\n", azColName[i], argv[i] ? argv[i] : "NULL");
  }
  printf("\n");
  return 0;
}