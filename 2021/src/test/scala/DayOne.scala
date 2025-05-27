val input = """199
200
208
210
200
207
241
269
260
263"""

class DayOneTestSuite extends munit.FunSuite {
  test("part one") {
    val expected = 7
    val result = DayOne.partOne(input)
    assertEquals(result, expected)
  }

  test("part two") {
    val expected = 5
    val result = DayOne.partTwo(input)
    assertEquals(result, expected)
  }
}
