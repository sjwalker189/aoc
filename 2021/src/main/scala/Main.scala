import scala.io

val Solutions: List[Solution] = List(
  DayOne
)

@main def main() = {

  println("-" * 32)
  println(" Advent of Code (2021) ")
  println("-" * 32)
  Solutions.foreach { solution =>
    val input = readFile(solution.source)
    println(" " + solution.title)
    println(s"  Part One: ${solution.partOne(input)}")
    println(s"  Part Two: ${solution.partTwo(input)}")
    println("-" * 32)
  }
}
