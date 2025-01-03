package zakumhelm

import "core:fmt"
import "core:math"

STAT_MIN :: 13
STAT_MAX :: 17
IMPORTANT_STATS :: 3

STAT_RANGE :: STAT_MAX - STAT_MIN + 1
TOTAL_HELMS :: STAT_RANGE * STAT_RANGE * STAT_RANGE * STAT_RANGE
SUM_MIN :: STAT_MIN*IMPORTANT_STATS
SUM_MAX :: STAT_MAX*IMPORTANT_STATS

is_helm_better :: proc(current_sum: int, str: int, dex: int, i: int, luk: int) -> bool {
    return (str+dex+luk) >= current_sum
}

count_better_helms :: proc(current_sum: int) -> (count: int) {
    for str in STAT_MIN..=STAT_MAX {
        for dex in STAT_MIN..=STAT_MAX {
            for i in STAT_MIN..=STAT_MAX {
                for luk in STAT_MIN..=STAT_MAX {
                    if is_helm_better(current_sum, str, dex, i, luk) do count += 1
                }
            }
        }   
    }

    return count
}

main :: proc() {
    fmt.println("    helms on floor")
    fmt.println("sum 1     2     3     4     5")

    for current_sum in SUM_MIN..=SUM_MAX {
        fmt.printf("%d  ", current_sum)

        n_better_helms := count_better_helms(current_sum)
        odds_better_helm := f64(n_better_helms) / f64(TOTAL_HELMS)

        for helms_on_floor in 1..=5 {
            odds_all_helms_worse := math.pow_f64(1-odds_better_helm, f64(helms_on_floor))
            fmt.printf("%.2f  ", 1-odds_all_helms_worse)
        }

        fmt.println("")
    }
}