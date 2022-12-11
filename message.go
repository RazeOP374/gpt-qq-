package main

import (
	chatgpt "GOproject/project1/a/gpt_token"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"strings"
)

func Cont(g *gin.Context) {
	Data := g.Request.Body
	read, _ := io.ReadAll(Data)
	postType := gjson.Get(string(read), "post_type").String()
	if postType == "message" {
		Message := gjson.Get(string(read), "message").String()
		if Message == "" {
			return
		} else {
			Chat := chatgpt.NewChat("eyJhbGciOiJkaXIiLCJlbmMiOiJBMjU2R0NNIn0..xeivJ6F-o9fQeNLj.7MXWqGO1jtVf_662byO50v8D9vpZcuUWoH80_xUkwyfAR9kI0-KTF9GwRghGeKTn9Xei2U-ct0mMkkbKuJAbfE7EuXDKAdHOKUlud3ys-EV3OMnUjcnPYDVvJ5NPDHEQ3c1HewYz_83PWGJrYegDbA1XxNbRNv7kryD_yJmqhasZ7vXzxC6IXJrDHrkBrNCQiIN4eDE62mcxbh4lHKqDQwadk2jesf2eGIh2yIMDfJW1Iibvrz7tPtddgea9vPxz_FvDSNV2vzwBRH0Y5mJ9sdsD3d9vCC7u9YTsn0zBDItaJiiXrv2Jr92V3-FLaIbfJuIvHZ20YVtG8DZ-3D46QaL6MurlQ0mbRjFpdgfmAs0ijjHMuYq2rVUGuzTtEtMa3Tq2-9RaYV5_yukTh1Mct1wvQLqYRN6HwiaBdqQKfW2bJG0AM5qQWKv4f5xe25O0cK4iR4bMypq8AiCalxe6wvJHAcgGBSrUw74H4kuJi4IbMgpfBazQbSf6qKprr3VwMDHGT95c-2Fi-NZyiMil0wNPc3dA9RBHP-QHLRrE1Qh2R0GF4RFmiW4jehplx9M0XSBKkmJ4ZkIrhzqYkzox42ney46jvZYdW37E3MhFzfxSEQz7LcOHUdwZoVAgzEc2b4DBR2ctahTbjpnRRNVHYs1y3aHI-Y-shaisHiNbssAxqQL9s6Eejw94DMXbzRW_ikG9AXzGxYNeNwf84bcYjN1EigarGfpWrW1b4TGwVULx-FZWZvVA5s2cuVU4Db-DsmlqBSdlW9kOgVRo4b_ZlQagMGovzW85iHelAtt3UKOp7ME_0dS4BzdqzMbDZexmlW4fPpNhX8OJix9Z0KCyO-jYCMdp9G70i_jRQ6eYLiA_0qgI7avPV4q7TFVFPzTsZ8KLhjvuCJ6ZsN0nwpCo-1WcvwPs6hulGuwuIo9Kz7NwCTAgS4eEq2VOqZFaq9KJBoNSagV6ZbX_p31ikVEDrCSOMnHhA3ADu_eFE1Se7F6t3JSkzVC0y-gD7hmm-sUMiyFupIVMj-g2S83stv9__BeXVuFluRv0YqGi8rx-tORvII44_q4VLuFINKvjY2jPlSRR25fgDxq8ewMjDrwR3cHjh3uqJtLPfamHYacFFwU-qOgTwUZH00dXDghncmk8yhDkcGGdKWFRxN0EIuovHRestkxlT2EMCeqJyfV-O-Vg9h7GN_15AjIEnWXmBYkOS9AYCPEt1kivuLF6B4cWxaHrKIBxCRzY0MOpZXAgqk7SjIzHyOrZ6uCV69KgpWxTnXFgfJhzJ_gUNpiUPMmkbivywb5SBEK59heQ3qq3P-5b9TXasOD9HEgVxJK1V3f-vcgm4s2mLbrppnNS5gmkV6e2tkpMnuIIlnSWjy1aEfd0G-0BazKVjLniI3FPBO-sexX5m_b9EG3wyzn7AqbINY93Z41ruzCWMwhgfzWG7LU1fv4Pp64zhL6dCNHQRxnMa_SqmUNX_aaTV6vA8LTdzYRv7xofFDP5EdRZhBNoVCGD3vE9nkz2kFSDxvSCA1HrkYZutPW1g9BrSXR45m5lrd-tHc-U6QDaeAuQmKX_9aU8uUBZu7JVvf1hdQkSBvQUgltzQ2dMdHyxJaRsQ_OtSJSsRYNCHUMKerNKakMTC9CTs-pTyD1ccb4gJahUVszC6rGQVzFBeEHldZiJWmbbmuXbr0cY0eMMqcdgbDl-W5Zdch6aRqt-gvvZNLES2WKjx8LFE5nXA0626RofyiKu-ue4rkwY4LB_ka_-3MKLvMNqHP5FIeTphpFzJs2J7WMmsKDIUT7lBno6cvXB5pQhxGF10Qa3YTVt2MOCUWntjqpejWQSn9NNABpWR7tOFpVh6DuUlkGUeb0BX-JiVq1NH-re-HLQx537jsq3sfrFE4mrBakjIuLzANoxX4oXXYrRKAtwMaK1JKxk8xV3f2MMOxkQKfYiIxOBU2pEB8NMN84t2n_O7jKdBE7OfzsiAmK7pzGu7UUHLHDkWp6FL97E8DOWHnzqcdQQm8M2hGALRl2cm1kbW7Oq4tWpZxeJp0rxnMUJpJSBqYwXWZw7595SdrHAdHcSe3XM5okrnCNK6r7BhpXwWJ5DPopqPgSpJsrteB6CClOS3FbVfEuSFb9GXBFlHi7EN63O5fCb9Uv_NLhj2gIo_JMNJThVjDC0XLTiSU6pK6_ExZRSjMxOpiNplxd3vAPpEe-zR1EMKpQnRDLE6ByVd1Y2b1JznH24I0dNYvICAHFY_bpD56iVbvawwtHKYp-WpPpBmg.0ZnTzoHUTvDcdV6h6G8zBg")
			res, err := Chat.Send(Message)
			if err != nil {
				return
			}
			ms, err := json.Marshal(res.Message.Content.Parts)
			if err != nil {
				return
			}
			fmt.Println(string(ms))
			text := string(ms)
			text = strings.Replace(text, "[", "", -1)
			text = strings.Replace(text, "\"", "", -1)
			text = strings.Replace(text, "]", "", -1)
			text = strings.Replace(text, "\n", "", -1)
			text = strings.Replace(text, "\n\n", "", -1)

			g.JSON(http.StatusOK, gin.H{
				"reply": text,
			})
			//fmt.Println(res.Message.Content.Parts)
		}
	}
}
