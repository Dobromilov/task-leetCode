func judgeCircle(moves string) bool {
    x:=0
    y:=0
    for i:=0; len(moves)>i; i++{
        if moves[i]=='U'{
            y++
        } else if moves[i]=='D'{
            y--
        } else if moves[i]=='L'{
            x--
        } else if moves[i]=='R'{
            x++
        }
    }
    return (x==0 && y==0)
}