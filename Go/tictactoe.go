/* * * * * * * * * * * * * * * * * * * * *
 *     Experimenting with Go             *
 * Yes, this is a sloppy program.        *
 * * * * * * * * * * * * * * * * * * * * */
package main
import (
    "fmt"
    "os"
)

// Our Tic Tac Toe board
var board[3][3]string

// initializes board to empty
func initBoard() {
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            board[j][i] = "-"
        }
    }
}

// win condition
func isOver() bool {
    var pieceArray[2]string
    pieceArray[0] = "x"
    pieceArray[1] = "o"
    for i := 0; i < 2; i++ {
    if      (board[0][0] == pieceArray[i] && board[0][1] == pieceArray[i] && board[0][2] == pieceArray[i]) ||
            (board[1][0] == pieceArray[i] && board[1][1] == pieceArray[i] && board[1][2] == pieceArray[i]) ||
            (board[2][0] == pieceArray[i] && board[2][1] == pieceArray[i] && board[2][2] == pieceArray[i]) ||
            (board[0][0] == pieceArray[i] && board[1][0] == pieceArray[i] && board[2][0] == pieceArray[i]) ||
            (board[0][1] == pieceArray[i] && board[1][1] == pieceArray[i] && board[2][1] == pieceArray[i]) ||
            (board[0][2] == pieceArray[i] && board[1][2] == pieceArray[i] && board[2][2] == pieceArray[i]) ||
            (board[0][0] == pieceArray[i] && board[1][1] == pieceArray[i] && board[2][2] == pieceArray[i]) ||
            (board[0][2] == pieceArray[i] && board[1][1] == pieceArray[i] && board[2][0] == pieceArray[i]) {
                return true
            }
    }
    return false
}

// prints out our board for us
func printBoard() {
    fmt.Println("   1  2  3")
    for i := 0; i < 3; i++ {
        fmt.Print(i+1, "  ")
        for j := 0; j < 3; j++ {
            fmt.Print(board[j][i], "  ")
        }
        fmt.Println("\n")
    }
}

// place piece into our board
func placePiece(x int, y int, piece string) {
    board[x][y] = piece;
}

// get input from command line
func getUserMove() {
    var x, y int
    fmt.Print("Place your piece: ")
    fmt.Scanf("%d", &x)
    fmt.Scanf("%d", &y)

    // adjust for array
    x -= 1
    y -= 1
    placePiece(x,y,"x")
}

func getComputerMove() {
    fmt.Println("I'm lazy")
}

func main() {
    initBoard()
    for i := 0; i < 9; i++ {
        printBoard()
        getUserMove()
        if isOver() {
            printBoard()
            fmt.Println("That's All folks!")
            os.Exit(1)
            break;
        }
        getComputerMove()
    }
    fmt.Println("You ended in a tie!")
}
