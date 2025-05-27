object DayOne extends Solution {
  val title = "Day One"
  val source = "dayone.txt"

  def partOne(input: String): Int = {
    countAsscendingValues(parseInput(input))
  }

  def partTwo(input: String): Int = {
    countAsscendingValues(
      parseInput(input)
        .sliding(3)
        .map(list => list.foldLeft(0)((acc, i) => acc + i))
        .toList
    )
  }

  // Count occurences of adjacent asscending values
  private def countAsscendingValues(list: List[Int]): Int = {
    list.zip(list.tail).count { case (a, b) =>
      a < b
    }
  }

  private def parseInput(raw: String): List[Int] = {
    return raw
      .split("\n")
      .filter(_.nonEmpty)
      .map(_.toInt)
      .toList
  }
}
