/*

 */
package entity

type Asset struct {
	Org       string
	TID      string
	Content   []byte
	Timestamp int
	Signature string
}

type Bundle interface{
	Category 
	Units		Items	
	Hash		[]byte
}

type Unit struct {
	UID
	NFT	[]byte
	CreationStamp	int
	Timestamp	int
	Signature	[]byte
	
}

func Create_Asset() (*Asset, error) {
	
}

//update_asset is a method of the Asset struct in order to provide new data, based on the past data
//This funtion takes in webtokens, verifies, and provides those resources to the asset struct
func (a *Asset) update_asset() (*Asset, error) {


}

func (a *Asset) New_Asset() {

}
