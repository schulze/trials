#include "gtest/gtest.h"

int square(int);

TEST(SquareTest, PositiveNos) { 
    EXPECT_EQ(9, square(3));
    EXPECT_EQ(4, square(2));
}

TEST (SquareTest, ZeroAndNegativeNos) { 
    ASSERT_EQ(0, square(0));
    ASSERT_EQ(9, square(-3));
}

int main(int argc, char **argv) {
  ::testing::InitGoogleTest(&argc, argv);
  return RUN_ALL_TESTS();
}
