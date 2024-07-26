/*
 * GMO Aozora Net Bank Open API
 *
 * <p>オープンAPI仕様書（PDF版）は下記リンクをご参照ください</p> <div>   <div style='display:inline-block;'><a style='text-decoration:none; font-weight:bold; color:#00b8d4;' href='https://gmo-aozora.com/business/service/api-specification.html' target='_blank'>オープンAPI仕様書</a></div><div style='display:inline-block; margin-left:2px; left:2px; width:10px; height:10px; border-top:2px solid #00b8d4; border-right:2px solid #00b8d4; transparent;-webkit-transform:rotate(45deg); transform: rotate(45deg);'></div> </div> <h4 style='margin-top:30px; border-left: solid 4px #1B2F48; padding: 0.1em 0.5em; color:#1B2F48;'>共通仕様</h4> <div style='width:100%; margin:10px;'>   <p style='font-weight:bold; color:#616161;'>＜HTTPリクエストヘッダ＞</p>   <div style='display:table; margin-left:10px; background-color:#29659b;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff;'>項目</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; color:#fff;'>仕様</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>プロトコル</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>HTTP1.1/HTTPS</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>charset</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>UTF-8</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>content-type</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>application/json</div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>domain_name</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>       本番環境：api.gmo-aozora.com</br>       開発環境：stg-api.gmo-aozora.com     </div>   </div>   <div style='display:table; margin-left:10px;'>     <div style='display:table-cell; min-width:130px; padding:9px; border:1px solid #fff; color:#fff; background-color:#29659b;'>メインURL</div>     <div style='display:table-cell; width:85%; padding:9px; border:1px solid #fff; background-color:#f8f8f8;'>       https://{domain_name}/ganb/api/corporation/{version}</br>       <span style='border-bottom:solid 1px;'>Version:1.x.x</span> の場合</br>       　https://api.gmo-aozora.com/ganb/api/corporation/<span style='border-bottom:solid 1px;'>v1</span>     </div>   </div> </div> <div style='margin:20px 10px;'>   <p style='font-weight:bold; color:#616161;'>＜リクエスト共通仕様＞</p>   <p style='padding-left:20px; font-weight:bold; color:#616161;'>NULLデータの扱い</p>   <p style='padding-left:40px;'>パラメータの値が空の場合、またはパラメータ自体が設定されていない場合、どちらもNULLとして扱います</p> </div> <div style='margin:20px 10px;'>   <p style='font-weight:bold; color:#616161;'>＜レスポンス共通仕様＞</p>   <p style='padding-left:20px; font-weight:bold; color:#616161;'>NULLデータの扱い</p>   <ul>     <li>レスポンスデータ</li>       <ul>         <li style='list-style-type:none;'>レスポンスデータの値が空の場合または、レスポンスデータ自体が設定されない場合は「項目自体を設定しません」と記載</li>       </ul>     <li>配列</li>       <ul>         <li style='list-style-type:none;'>配列の要素の値が空の場合は「空のリスト」と記載</li>         <li style='list-style-type:none;'>配列自体が設定されない場合は「項目自体を設定しません」と記載</li>       </ul>   </ul> </div> <div style='margin:20px 10px;'>   <p style='font-weight:bold; color:#616161;'>＜更新系APIに関する注意事項＞</p>   <ul>     <li style='list-style-type:none;'>更新系処理がタイムアウトとなった場合、処理自体は実行されている可能性がありますので、</li>     <li style='list-style-type:none;'>再実行を行う必要がある場合は必ず照会系の処理で実行状況を確認してから再実行を行ってください</li>   </ul> </div>
 *
 * API version: 1.1.12
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Balance struct {
	// 口座ID 半角英数字 口座を識別するID
	AccountId string `json:"accountId"`
	// 科目コード 半角数字 ・01=普通預金（有利息） ・02=普通預金（決済用） ・11=円定期預金 ・51=外貨普通預金 ・81=証券コネクト口座
	AccountTypeCode string `json:"accountTypeCode"`
	// 科目名 全角文字 科目コードの名称
	AccountTypeName string `json:"accountTypeName"`
	// 現在残高 半角数字　マイナス含む 基準日時における現在残高 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01=普通預金（有利息） ・02=普通預金（決済用） ・11=円定期預金 ・81=証券コネクト口座
	Balance string `json:"balance,omitempty"`
	// 基準日 半角文字 残高および引出可能額を照会した基準日を示します YYYY-MM-DD形式
	BaseDate string `json:"baseDate"`
	// 基準時刻 半角文字 残高および引出可能額を照会した基準時刻を示します HH:MM:SS+09:00形式
	BaseTime string `json:"baseTime"`
	// 支払可能残高 半角数字　マイナス含む 応答時点での引出可能額を示します 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01=普通預金（有利息） ・02=普通預金（決済用）
	WithdrawableAmount string `json:"withdrawableAmount,omitempty"`
	// 前日残高 半角数字　マイナス含む 日付が変わった直後は、銀行の締め処理が終わるまでは前々日の残高となります 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01=普通預金（有利息） ・02=普通預金（決済用）
	PreviousDayBalance string `json:"previousDayBalance,omitempty"`
	// 前月末残高 半角数字　マイナス含む 月が変わった直後は、銀行の締め処理が終わるまでは前々月の残高となります  科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・01=普通預金（有利息） ・02=普通預金（決済用）
	PreviousMonthBalance string `json:"previousMonthBalance,omitempty"`
	// 通貨コード 半角文字 ISO4217に準拠した通貨コード
	CurrencyCode string `json:"currencyCode"`
	// 通貨名 全角文字 ISO4217に準拠した通貨コードの当行での名称
	CurrencyName string `json:"currencyName"`
	// 外貨残高 半角数字　少数点および小数部最大3桁　マイナス含む 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・51=外貨普通預金
	FcyTotalBalance string `json:"fcyTotalBalance,omitempty"`
	// 外貨レート 半角数字　少数点および小数部最大3桁　マイナス含む 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・51=外貨普通預金
	Ttb string `json:"ttb,omitempty"`
	// 外貨レート基準日 半角文字 外貨レートの基準日を示します YYYY-MM-DD形式 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・51=外貨普通預金
	BaseRateDate string `json:"baseRateDate,omitempty"`
	// 外貨レート基準時刻 半角文字 外貨レートの基準時刻を示します HH:MM:SS+09:00形式 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・51=外貨普通預金
	BaseRateTime string `json:"baseRateTime,omitempty"`
	// 外貨円換算額 半角数字　マイナス含む 外貨残高を円に換算した額 科目コードが以下の場合のみ設定されます 該当しない場合は項目自体を設定しません ・51=外貨普通預金
	YenEquivalent string `json:"yenEquivalent,omitempty"`
}
