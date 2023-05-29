package q1900

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 将 MovieRentingSystem 对象用 n 个商店和 entries 表示的电影列表初始化。
type MovieRentingSystem struct {
	// shopi moviei price rented
	shop   map[int]map[int]movie
	movie  map[int][]movie
	rented map[string]int
}

type movie struct {
	shopi  int
	moviei int
	price  int
}

func Q1912Constructor(n int, entries [][]int) MovieRentingSystem {
	shop := make(map[int]map[int]movie)
	movies := make(map[int][]movie)
	for i := 0; i < len(entries); i++ {
		m := movie{shopi: entries[i][0], moviei: entries[i][1], price: entries[i][2]}
		if _, ok := shop[entries[i][0]]; !ok {
			shop[m.shopi] = make(map[int]movie)
		}
		shop[m.shopi][m.moviei] = m
		movies[m.moviei] = append(movies[m.moviei], m)
	}
	for k, ms := range movies {
		sort.Slice(ms, func(i, j int) bool {
			if ms[i].price == ms[j].price {
				return ms[i].shopi < ms[j].shopi
			}
			return ms[i].price < ms[j].price
		})
		movies[k] = ms
	}
	return MovieRentingSystem{shop: shop, movie: movies, rented: make(map[string]int)}
}

// Search 找到拥有指定电影且 **未借出** 的商店中 最便宜的 5 个 。
//商店需要按照价格升序排序，如果价格相同，则 shopi 较小的商店排在前面。
//如果查询结果少于 5 个商店，
//则将它们全部返回。如果查询结果没有任何商店，则返回空列表。
//
// select * from movie_renting_system where movie = ? and rented = 0 order by price, shop limit 5;
func (this *MovieRentingSystem) Search(movie int) []int {
	if ms := this.movie[movie]; len(ms) > 0 {
		r := make([]int, 0, 5)
		for i := 0; i < len(ms) && len(r) < 5; i++ {
			if this.rented[fmt.Sprintf("%d_%d", ms[i].shopi, ms[i].moviei)] <= 0 {
				r = append(r, ms[i].shopi)
			}
		}
		return r
	}
	return []int{}
}

// Rent 从指定商店借出指定电影，题目保证指定电影在指定商店未借出 。
//
// update movie_renting_system set rented = 1 where shop = ? and movie = ?;
func (this *MovieRentingSystem) Rent(shop int, movie int) {
	if s, ok := this.shop[shop]; ok {
		if m, ok := s[movie]; ok {
			this.rented[fmt.Sprintf("%d_%d", shop, movie)] = m.price
		}
	}
}

// Drop 在指定商店返还之前已借出的指定电影。
// update movie_renting_system set rented = 0 where shop = ? and movie = ?;
func (this *MovieRentingSystem) Drop(shop int, movie int) {
	delete(this.rented, fmt.Sprintf("%d_%d", shop, movie))
}

// Report 返回 最便宜的 5 部已借出电影 （可能有重复的电影 ID），将结果用二维列表 res 返回，
//其中 `res[j] = [shopj, moviej]` 表示第 j 便宜的已借出电影是从商店 shopj 借出的电影 moviej 。
//res 中的电影需要按 价格 升序排序；如果价格相同，则 shopj 较小 的排在前面；
//如果仍然相同，则 moviej 较小 的排在前面。如果当前借出的电影小于 5 部，则将它们全部返回。
//如果当前没有借出电影，则返回一个空的列表。
//
// select * from movie_renting_system where rented = 1 order by price, shop, movie limit 5;
func (this *MovieRentingSystem) Report() [][]int {
	values := make([]movie, 0, len(this.rented))
	for k, v := range this.rented {
		if v > 0 {
			ss := strings.Split(k, "_")
			shop, _ := strconv.Atoi(ss[0])
			m, _ := strconv.Atoi(ss[1])
			values = append(values, movie{shopi: shop, moviei: m, price: v})
		}
	}
	sort.Slice(values, func(i, j int) bool {
		if values[i].price == values[j].price {
			if values[i].shopi == values[j].shopi {
				return values[i].moviei < values[j].moviei
			}
			return values[i].shopi < values[j].shopi
		}
		return values[i].price < values[j].price
	})
	r := [][]int{}
	for i := 0; i < len(values) && len(r) < 5; i++ {
		r = append(r, []int{values[i].shopi, values[i].moviei})
	}
	return r
}
