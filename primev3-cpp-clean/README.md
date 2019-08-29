### pre-require
1. boost, need to be instal manually
2. cmake, need to be installed manually
3. gtest (will be installed automatically)

### build, test and run
$ cd primev3-cpp-clean
$ mkdir release
$ cd release 
$ cmake -DCMAKE_BUILD_TYPE=Release ..
$ make all test
$ cd .. && release/pokercpp

### Scope
Leverage State of the art C++11 implements Poker Algorithm
使用最新的C++工程技术实现poker算法，比较下C++和golang的性能效率

### Part 1: build a project for C++ with gtest/boost/cmake

### Part 2: poker algorithm implmentation

### Part 3: go calls c++ library and vice visa

cmake with gtest

1. install cmake
2. https://github.com/google/googletest/blob/master/googletest/README.md, build with CMake.  Use CMake to download GoogleTest as part of the build's configure step. This is just a little more complex, but doesn't have the limitations of the other methods.
The last of the above methods is implemented with a small piece of CMake code in a separate file (e.g. CMakeLists.txt.in) which is copied to the build area and then invoked as a sub-build during the CMake stage. That directory is then pulled into the main build with add_subdirectory(). 
New file CMakeLists.txt.in:

cmake_minimum_required(VERSION 2.8.2)

project(googletest-download NONE)

include(ExternalProject)
ExternalProject_Add(googletest
  GIT_REPOSITORY    https://github.com/google/googletest.git
  GIT_TAG           master
  SOURCE_DIR        "${CMAKE_CURRENT_BINARY_DIR}/googletest-src"
  BINARY_DIR        "${CMAKE_CURRENT_BINARY_DIR}/googletest-build"
  CONFIGURE_COMMAND ""
  BUILD_COMMAND     ""
  INSTALL_COMMAND   ""
  TEST_COMMAND      ""
)
# Download and unpack googletest at configure time
configure_file(CMakeLists.txt.in googletest-download/CMakeLists.txt)
execute_process(COMMAND ${CMAKE_COMMAND} -G "${CMAKE_GENERATOR}" .
  RESULT_VARIABLE result
  WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}/googletest-download )
if(result)
  message(FATAL_ERROR "CMake step for googletest failed: ${result}")
endif()
execute_process(COMMAND ${CMAKE_COMMAND} --build .
  RESULT_VARIABLE result
  WORKING_DIRECTORY ${CMAKE_CURRENT_BINARY_DIR}/googletest-download )
if(result)
  message(FATAL_ERROR "Build step for googletest failed: ${result}")
endif()

# Prevent overriding the parent project's compiler/linker
# settings on Windows
set(gtest_force_shared_crt ON CACHE BOOL "" FORCE)

# Add googletest directly to our build. This defines
# the gtest and gtest_main targets.
add_subdirectory(${CMAKE_CURRENT_BINARY_DIR}/googletest-src
                 ${CMAKE_CURRENT_BINARY_DIR}/googletest-build
                 EXCLUDE_FROM_ALL)

# The gtest/gtest_main targets carry header search path
# dependencies automatically when using CMake 2.8.11 or
# later. Otherwise we have to add them here ourselves.
if (CMAKE_VERSION VERSION_LESS 2.8.11)
  include_directories("${gtest_SOURCE_DIR}/include")
endif()

# Now simply link against gtest or gtest_main as needed. Eg
add_executable(example example.cpp)
target_link_libraries(example gtest_main)
add_test(NAME example_test COMMAND example)


3. run ctest -V or make all test
