# 设置要求的最低CMake版本
cmake_minimum_required(VERSION 3.20)

# 设置项目名称、版本所用语言
project(read_csv VERSION 1.0.0 LANGUAGES C)

# 设置C语言标准
set(CMAKE_C_STANDARD 17)

# 设置可执行文件的生成路径
set(EXECUTABLE_OUTPUT_PATH ${PROJECT_SOURCE_DIR}/bin)

# 设置可执行文件的生成
add_executable(read_csv
               src/main.c
               src/tools.c)
