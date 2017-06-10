object Main {
    def main(args: Array[String]) {
        val list = scala.io.StdIn.readLine.split(" ").map(_.toInt)
        val a = list(0)
        val b = list(1)
        println(a + b)
    }
}
