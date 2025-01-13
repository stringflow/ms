package rudolph

import "core:fmt"
import "core:math"
import "core:slice"

PRESENTS :: 838
ATTEMPTS :: PRESENTS / 300
CURRENT_SUM :: 21

STAT_MIN :: -4
STAT_MAX :: 6
STAT_RANGE :: STAT_MAX - STAT_MIN + 1

ACC_MIN :: 15
ACC_MAX :: 25
ACC_RANGE :: ACC_MAX - ACC_MIN + 1

TOTAL_NOSES :: STAT_RANGE * STAT_RANGE * ACC_RANGE

is_better_nose :: proc(target: int, acc: int, str: int, dex: int, i: int, luk: int) -> bool {
    str := max(str, 0)
    dex := max(dex, 0)
    i := max(i, 0)
    luk := max(luk, 0)
    return int(f64(acc) + f64(dex)*0.8 + f64(luk)*0.5) >= target
}

count_better_noses :: proc(target: int) -> (count: int) {
    for acc in ACC_MIN..=ACC_MAX {
        for dex in STAT_MIN..=STAT_MAX {
            for luk in STAT_MIN..=STAT_MAX {
                if is_better_nose(target, acc, 0, dex, 0, luk) do count += 1
            }
        }
    }

    return count
}

main :: proc() {
    fmt.printfln("current nose: %v", CURRENT_SUM)
    fmt.printfln("%d presents (%d attempts)", PRESENTS, ATTEMPTS)

    better_noses := count_better_noses(CURRENT_SUM)
    odds_better_nose := f64(better_noses) / f64(TOTAL_NOSES)
    fmt.printfln("odds of a better nose: %.2f%%", odds_better_nose*100)

    odds_all_noses_worse := math.pow_f64(1-odds_better_nose, ATTEMPTS)
    fmt.printfln("odds of a better nose in %v attempts: %.2f%%", ATTEMPTS, (1-odds_all_noses_worse)*100)
}