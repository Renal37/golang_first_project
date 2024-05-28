package main

import "fmt"

func main() {
    allNewsPosts := []string{
        "Post 1",
        "Post 2",
        "Post 3",
        "Post 4",
        "Post 5",
    }
    
    fmt.Println("<mainpage>")
    showMainPage(allNewsPosts[:3:3])
    fmt.Println("</mainpage>")

    fmt.Println("<all>")
    showPosts(allNewsPosts)
    fmt.Println("</all>")

}

func showMainPage(posts []string){
	postsWithAds := append(posts,"Click here")

	showPosts(postsWithAds)
}
func showPosts(posts []string){
	for _, post := range posts{
		fmt.Println(post)
	}
}