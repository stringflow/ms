package ms

import "core:fmt"
import "core:math"

STAT_MIN :: 13
STAT_MAX :: 17
IMPORTANT_STATS :: 3

SUM_MIN :: STAT_MIN*IMPORTANT_STATS
SUM_MAX :: STAT_MAX*IMPORTANT_STATS

count_better_helms :: proc(target_sum: int, stats_left: int, current_sum: int) -> int {
    if stats_left == 0 do return current_sum >= target_sum ? 1 : 0
    
    count := 0
    for stat in STAT_MIN..=STAT_MAX {
        count += count_better_helms(target_sum, stats_left - 1, current_sum + stat)
    }

    return count
}

main :: proc() {
    fmt.println("    helms on floor")
    fmt.println("sum 1     2     3     4     5")

    total_possible_helms := int(math.pow_f64(STAT_MAX - STAT_MIN + 1, IMPORTANT_STATS))

    for current_sum in SUM_MIN..=SUM_MAX {
        fmt.printf("%d  ", current_sum)

        n_better_helms := count_better_helms(current_sum, IMPORTANT_STATS, 0)
        odds_better_helm := f64(n_better_helms) / f64(total_possible_helms)

        for helms_on_floor in 1..=5 {
            odds_all_helms_worse := math.pow_f64(1-odds_better_helm, f64(helms_on_floor))
            fmt.printf("%.2f  ", 1-odds_all_helms_worse)
        }

        fmt.println("")
    }
}