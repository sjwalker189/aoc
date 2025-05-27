import java.nio.file

trait Solution {
  val title: String
  val source: String
  def partOne(input: String): Int
  def partTwo(input: String): Int
}

def readFile(filename: String): String = {
  val cwd = file.Paths.get("")
  val path = file.Paths.get(cwd.toString(), "resources", filename)
  io.Source.fromFile(path.toString(), "utf-8").mkString
}
