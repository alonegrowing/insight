package macro


type Comment struct {
	ID int64 `json:"id"`
	Content string `json:"content"`
	Created string `json:"created"`
	Member Member `json:"member"`
	Interactive Interactive `json:"interactive"`
}

var Comments = []Comment {
	{
		ID: 1,
		Content: "如果你爱过一个人，你定能明白我现在的心情",
		Created: "2021-09-21",
		Member: Member{
			ID:     351936818,
			Name:   "Twice-Chic",
			Avatar: "https://tvax2.sinaimg.cn/crop.0.0.1080.1080.1024/78d674f3ly8gdrjgfmd1zj20u00u0ack.jpg",
		},
		Interactive: Interactive{
			GoodNum:    216,
			CommentNum: 2,
			ShareNum:   1122,
		},
	},
	{
		ID: 2,
		Content: "既有少女的清甜，又兼具美少年气质! 出道13了好难得",
		Created: "2021-09-21",
		Member: Member{
			ID:     351936818,
			Name:   "敏儿好学YA",
			Avatar: "https://tvax3.sinaimg.cn/crop.0.0.996.996.1024/006YiKGDly8glwu66k5jtj30ro0rodi2.jpg",
		},
		Interactive: Interactive{
			GoodNum:    216,
			CommentNum: 2,
			ShareNum:   1122,
		},
	},
	{
		ID: 3,
		Content: "少管明星的那点鸡毛蒜皮的小事，多管一管自己的事",
		Created: "2021-09-21",
		Member: Member{
			ID:     351936818,
			Name:   "刷神彡勋王",
			Avatar: "https://tvax3.sinaimg.cn/crop.0.0.1080.1080.1024/0061icylly8gm45u4hzcij30u00u0gni.jpg",
		},
		Interactive: Interactive{
			GoodNum:    216,
			CommentNum: 2,
			ShareNum:   1122,
		},
	},
	{
		ID: 4,
		Content: "没有以前帅了[doge]",
		Created: "2021-09-21",
		Member: Member{
			ID:     351936818,
			Name:   "送我一只酷盖A",
			Avatar: "https://tvax2.sinaimg.cn/crop.0.0.1080.1080.1024/006pHLWgly8gkdyigzlr7j30u00u0acc.jpg",
		},
		Interactive: Interactive{
			GoodNum:    126,
			CommentNum: 2,
			ShareNum:   122,
		},
	},
	{
		ID: 5,
		Content: "跟他姐姐好像，可是他姐姐还行啊，他长得就好显老",
		Created: "2021-09-21",
		Member: Member{
			ID:     351936818,
			Name:   "棕榈君的备忘录",
			Avatar: "https://tvax2.sinaimg.cn/crop.0.0.1080.1080.1024/0085SNm6ly8gepjkyiscvj30u00u00vg.jpg",
		},
		Interactive: Interactive{
			GoodNum:    216,
			CommentNum: 2,
			ShareNum:   1122,
		},
	},
}