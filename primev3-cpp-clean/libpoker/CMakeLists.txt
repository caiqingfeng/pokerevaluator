enable_testing()
include_directories("${gtest_SOURCE_DIR}/include")

add_library(poker LYCardHelpers.cpp LYScoreMap.cpp LYOffset.cpp)

# depending on the framework, you need to link to it
# target_link_libraries(pokercpp_test gtest_main)

# now register the executable with CTest
# add_test(NAME pokercpp_test COMMAND pokercpp_test)
