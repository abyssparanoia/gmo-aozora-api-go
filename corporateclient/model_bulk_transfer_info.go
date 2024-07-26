/*
 * GMO Aozora Net Bank Open API
 *
 * <p>オープンAPI仕様書（PDF版）は下記リンクをご参照ください</p> <div>   <div style='display:inline-block;'><a style='text-decoration:none; font-weight:bold; color:#00b8d4;' href='https://gmo-aozora.com/business/service/api-specification.html' target='_blank'>オープンAPI仕様書</a></div><div style='display:inline-block; margin-left:2px; left:2px; width:10px; height:10px; border-top:2px solid #00b8d4; border-right:2px solid #00b8d4; transparent;-webkit-transform:rotate(45deg); transform: rotate(45deg);'></div> </div> <h4 style='margin-top:30px; border-left: solid 4px #1B2F48; padding: 0.1em 0.5em; color:#1B2F48;'>共通仕様</h4> <div style='width:100%; margin:10px;'>   <p style='font-weight:bold; color:#616161;'>＜HTTPリクエストヘッダ＞</p>   <div style='display:table; margin-left:10px; background-color:#29659b;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff;'>項目</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; color:#fff;'>仕様</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>プロトコル</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>HTTP1.1/HTTPS</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>charset</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>UTF-8</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>content-type</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>application/json</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>domain_name</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>       本番環境：api.gmo-aozora.com</br>       開発環境：stg-api.gmo-aozora.com     </div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>メインURL</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>       https://{domain_name}/ganb/api/corporation/{version}</br>       <span style='border-bottom:solid 1px;'>Version:1.x.x</span> の場合</br>       　https://api.gmo-aozora.com/ganb/api/corporation/<span style='border-bottom:solid 1px;'>v1</span>     </div>   </div> </div> <div style='margin:20px 10px;'>   <p style='font-weight:bold; color:#616161;'>＜リクエスト共通仕様＞</p>   <p style='padding-left:20px; font-weight:bold; color:#616161;'>NULLデータの扱い</p>   <p style='padding-left:40px;'>パラメータの値が空の場合、またはパラメータ自体が設定されていない場合、どちらもNULLとして扱います</p> </div> <div style='margin:20px 10px;'>   <p style='font-weight:bold; color:#616161;'>＜レスポンス共通仕様＞</p>   <p style='padding-left:20px; font-weight:bold; color:#616161;'>NULLデータの扱い</p>   <ul>     <li>レスポンスデータ</li>       <ul>         <li style='list-style-type:none;'>レスポンスデータの値が空の場合または、レスポンスデータ自体が設定されない場合は「項目自体を設定しません」と記載</li>       </ul>     <li>配列</li>       <ul>         <li style='list-style-type:none;'>配列の要素の値が空の場合は「空のリスト」と記載</li>         <li style='list-style-type:none;'>配列自体が設定されない場合は「項目自体を設定しません」と記載</li>       </ul>   </ul> </div> <div style='margin:20px 10px;'>   <p style='font-weight:bold; color:#616161;'>＜更新系APIに関する注意事項＞</p>   <ul>     <li style='list-style-type:none;'>更新系処理がタイムアウトとなった場合、処理自体は実行されている可能性がありますので、</li>     <li style='list-style-type:none;'>再実行を行う必要がある場合は必ず照会系の処理で実行状況を確認してから再実行を行ってください</li>   </ul> </div> 
 *
 * API version: 1.1.12
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type BulkTransferInfo struct {
	// 明細番号 半角数字 重複/0はエラー　1～9999とします 
	ItemId string `json:"itemId,omitempty"`
	// 被仕向金融機関番号 半角数字 
	BeneficiaryBankCode string `json:"beneficiaryBankCode,omitempty"`
	// 被仕向金融機関名カナ 半角文字 該当する情報が無い場合は項目自体を設定しません 
	BeneficiaryBankName string `json:"beneficiaryBankName,omitempty"`
	// 被仕向支店番号 半角数字 
	BeneficiaryBranchCode string `json:"beneficiaryBranchCode,omitempty"`
	// 被仕向支店名カナ 半角文字 該当する情報が無い場合は項目自体を設定しません 
	BeneficiaryBranchName string `json:"beneficiaryBranchName,omitempty"`
	// 手形交換所番号 半角文字 該当する情報が無い場合は項目自体を設定しません 
	ClearingHouseName string `json:"clearingHouseName,omitempty"`
	// 科目コード（預金種別コード） 半角数字 1：普通、2：当座、4：貯蓄、9：その他 
	AccountTypeCode string `json:"accountTypeCode,omitempty"`
	// 口座番号 半角数字 7桁未満の番号は右詰で、前ゼロで埋めること 
	AccountNumber string `json:"accountNumber,omitempty"`
	// 受取人名 半角文字 該当する情報が無い場合は項目自体を設定しません 
	BeneficiaryName string `json:"beneficiaryName,omitempty"`
	// 振込金額 半角数字 1以上9,999,999,999円以下　数値のみでカンマ、マイナス不可 
	TransferAmount string `json:"transferAmount,omitempty"`
	// 新規コード 半角文字 該当する情報が無い場合は項目自体を設定しません 
	NewCode string `json:"newCode,omitempty"`
	// EDI情報 半角文字 該当する情報が無い場合は項目自体を設定しません 
	EdiInfo string `json:"ediInfo,omitempty"`
	// 振込指定区分 半角文字 該当する情報が無い場合は項目自体を設定しません 
	TransferDesignatedType string `json:"transferDesignatedType,omitempty"`
	// 識別表示 半角文字 該当する情報が無い場合は項目自体を設定しません 
	Identification string `json:"identification,omitempty"`
	// 振込明細結果 振込明細結果のリスト 正常時のみ応答 該当する情報が無い場合は空のリストを返却 
	TransferDetailResponses []TransferDetailResponse `json:"transferDetailResponses,omitempty"`
	// 不能明細情報 不能明細情報のリスト 該当する情報が無い場合は項目自体を設定しません 
	UnableDetailInfos []UnableDetailInfo `json:"unableDetailInfos,omitempty"`
}
