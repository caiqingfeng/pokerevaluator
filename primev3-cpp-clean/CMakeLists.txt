cmake_minimum_required (VERSION 2.8.2)
project (poker)
configure_file(CMakeLists.txt.in googletest-download/CMakeLists.txt)
set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
find_package(Boost REQUIRED)
if(Boost_FOUND)
  include_directories(${Boost_INCLUDE_DIRS})
endif()

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

add_subdirectory(libpoker)

enable_testing()
include_directories("${gtest_SOURCE_DIR}/include")

add_executable(pokercpp src/poker.cpp )
add_executable(pokercpp_test test/LYCardHelpers_tests.cpp 
        test/LYOffset_tests.cpp test/LYScoreMap_tests.cpp
        test/LYScoreMapHand_tests.cpp)

add_dependencies(pokercpp_test gtest)

# depending on the framework, you need to link to it
target_link_libraries(pokercpp_test gtest_main poker)
target_link_libraries(pokercpp poker)

# now register the executable with CTest
add_test(NAME pokercpp_test COMMAND pokercpp_test)
