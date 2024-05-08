
var __hasSSR__ = false;
if (__hasSSR__) {
timings.ssr = Date.now();
}
window.LZD_RETCODE_PAGENAME = 'pdp-m-revamp';
window.LZD_RETCODE_SAMPLE = 100;
window.g_config = window.g_config || {};
window.g_config.regionID = "TH";
window.g_config.language = "th";
var paths = {};
var maps = {
"ReactDOM": {
"root": "react-dom",
"umd": "//g.lazcdn.com/g/lzd/assets/1.2.13/react-dom/16.8.0/react-dom.production.min"
},
"React": {
"root": "react",
"umd": "//g.lazcdn.com/g/lzd/assets/1.2.13/react/16.8.0/react.production.min"
},
"Next": {
"root": "@alife/next",
"umd": "//g.lazcdn.com/g/lzd/assets/1.2.13/next/0.19.21/next.min"
},
"Env": {
"root": "@ali/lzd-h5-utils-env",
"umd": "//g.lazcdn.com/g/code/npm/@ali/lzd-h5-utils-env/1.5.12/index"
},
"Cookie": {
"root": "@ali/lzd-h5-utils-cookie",
"umd": "//g.lazcdn.com/g/code/npm/@ali/lzd-h5-utils-cookie/1.2.10/index"
},
"Logger": {
"root": "@ali/lzd-h5-utils-logger",
"umd": "//g.lazcdn.com/g/code/npm/@ali/lzd-h5-utils-logger/1.1.52/index"
},
"Mtop": {
"root": "@ali/lzd-h5-utils-mtop",
"umd": "//g.lazcdn.com/g/code/npm/@ali/lzd-h5-utils-mtop/1.2.56/index"
},
"Icon": {
"root": "@ali/lzd-h5-utils-icon",
"umd": "//g.lazcdn.com/g/code/npm/@ali/lzd-h5-utils-icon/1.0.8/index"
}
};
var key;
for (key in maps) {
if (maps.hasOwnProperty(key)) {
if (window[key]) {
define(maps[key].root, [], (function (key) {
return window[key];
})(key));
}
else {
paths[maps[key].root] = maps[key].umd;
}
}
}
requirejs.config({
paths: paths,
waitSeconds: 0
});
window.__i18n__ = {"th":{"pdp_static.addOnService.buyMoreGradient":"รับส่วนลด {discount} จากยอดซื้อ {total}","pdp_static.addOnService.title":"ซื้อพร้อมบริการเสริม","pdp_static.addOnService.title2":"กรุณาเลือกการบริการเสริม","pdp_static.addOnService.total":"ยอดรวมทั้งสิ้น:","pdp_static.addOnService.viewItem":"ดูสินค้า","pdp_static.add_on_service.title":"ซื้อพร้อมบริการเสริม","pdp_static.add_on_service.title.pc":"ซื้อพร้อมบริการเสริม","pdp_static.add_on_service.title2":"ซื้อพร้อมบริการเสริม","pdp_static.add_on_service.total":"ยอดรวมทั้งสิ้น","pdp_static.add_on_service.view_item":"ดูสินค้า","pdp_static.age-restrict.dontaskagain":"ไม่ต้องถามอีก","pdp_static.age-restrict.title":"ยืนยันอายุของท่าน","pdp_static.cart.addCart":"เพิ่มไปยังรถเข็น","pdp_static.cart.addWishlist":"เพิ่มสินค้านี้ลงรายการที่ชอบ","pdp_static.cart.add_cart":"เพิ่มไปยังรถเข็น","pdp_static.cart.add_wishlist":"เพิ่มสินค้านี้ลงรายการที่ชอบ","pdp_static.cart.alreadyInWishlist":"อยู่ในหน้ารายการที่ชอบแล้ว","pdp_static.cart.already_in_wishlist":"อยู่ในหน้ารายการที่ชอบแล้ว","pdp_static.cart.buy_tips":"สั่งซื้อได้อย่างไร?","pdp_static.cart.can_not_buy_tips":"สินค้านี้ไม่สามารถสั่งซื้อได้","pdp_static.cart.chooseOptions":"กรุณาเลือกตัวเลือก","pdp_static.cart.choose_options":"กรุณาเลือกตัวเลือก","pdp_static.cart.comfirm":"Confirm","pdp_static.cart.commingSoon":"จะมาใหม่เร็วๆนี้","pdp_static.cart.comming_soon":"จะมาใหม่เร็วๆนี้","pdp_static.cart.failedAddWishlist":"ไม่สามารถเพิ่มสินค้าไปยังรายการที่ชอบได้","pdp_static.cart.failedDelWishlist":"ไม่สามารถลบสินค้าออกจากรายการที่ชอบได้","pdp_static.cart.failedTip":"ไม่สามารถเพิ่มสินค้าไปรถเข็นได้","pdp_static.cart.failed_add_wishlist":"เพิ่มสินค้านี้ไปรายการที่ชอบไม่สำเร็จ","pdp_static.cart.failed_del_wishlist":"ไม่สามารถลบสินค้าชิ้นนี้ออกจากรายการที่ชอบได้","pdp_static.cart.failed_tip":"ไม่สามารถเพิ่มสินค้าไปรถเข็นได้","pdp_static.cart.goToCart":"ไปที่รถเข็น","pdp_static.cart.go_to_cart":"ไปที่รถเข็น","pdp_static.cart.successAddWishlist":"เพิ่มสินค้าไปยังรายการที่ชอบได้สำเร็จ","pdp_static.cart.successDelWishlist":"ลบสินค้าไปยังรายการที่ชอบได้สำเร็จ","pdp_static.cart.successTip":"เพิ่มสินค้าไปยังรถเข็นได้สำเร็จ","pdp_static.cart.success_add_wishlist":"เพิ่มสินค้านี้ไปรายการที่ชอบได้สำเร็จ","pdp_static.cart.success_del_wishlist":"ลบสินค้าชินนี้ออกจากรายการที่ชอบสำเร็จ","pdp_static.cart.success_tip":"เพิ่มสินค้าลงรถเข็นสำเร็จ","pdp_static.combo.addCartText":"เพิ่ม {count} {count, plural, one {ชิ้น} other {ชิ้น}} ไปยังรถเข็น","pdp_static.combo.add_cart_text":"เพิ่ม {count} {count, plural, one {ชิ้น} other {ชิ้น}} ไปยังรถเข็น","pdp_static.combo.comboCountText":"ชุดสุดคุ้มจำนวน {count} {count, plural, one {ชุด} other {ชุด}}","pdp_static.combo.comboOfferText":"ข้อเสนอสินค้าชุดสุดคุ้ม","pdp_static.combo.comboPrice":"ราคาสินค้าชุดสุดคุ้ม","pdp_static.combo.comboTotal":"ยอดรวมสินค้าชุดสุดคุ้ม","pdp_static.combo.combo_count_text":"{count} ชิ้น {count, plural, one {ข้อเสนอ}} other {ข้อเสนอ}}","pdp_static.combo.combo_offer_text":"ข้อเสนอสินค้าชุดสุดคุ้ม","pdp_static.combo.combo_price":"ราคาสินค้าชุดสุดคุ้ม","pdp_static.combo.combo_total":"ยอดรวมสินค้าชุดสุดคุ้ม","pdp_static.combo.discount":"ส่วนลด","pdp_static.combo.edit":"แก้ไข","pdp_static.combo.ellipsis":"…","pdp_static.combo.promotion":"โปรโมชั่น","pdp_static.combo.save":"ประหยัด","pdp_static.combo.save.test":"pdp_static.combo.save.test","pdp_static.combo.save.test2222":"testwr323r","pdp_static.combo.save.test888888888":"2.34234234E8","pdp_static.combo.saveUpTo":"ประหยัดสูงสุดถึง {saving}","pdp_static.combo.save_up_to":"ประหยัดถึง {saving}","pdp_static.combo.sliderIndexText":"{current} จาก {total}","pdp_static.combo.slider_index_text":"{current} จาก {total}","pdp_static.combo.subtotal":"ยอดรวม","pdp_static.combo.title":"สินค้านี้จะถูกซื้อคู่กันบ่อย","pdp_static.combo.total":"รวมทั้งสิ้น","pdp_static.comma":",","pdp_static.common.abuse_text":"รายงานความเห็นนี้","pdp_static.common.bottombar.gotocart":"ไปที่รถเข็น","pdp_static.common.brand":"แบรนด์","pdp_static.common.cancel":"ยกเลิก","pdp_static.common.comma":",","pdp_static.common.confirm":"ยืนยัน","pdp_static.common.darazHeader.searchBar":"Search in Daraz","pdp_static.common.got_it":"เข้าใจ","pdp_static.common.helpful":"ไม่เป็นประโยชน์","pdp_static.common.in":"ใน","pdp_static.common.login":"เข้าสู่ระบบ","pdp_static.common.more":"เพิ่มเติม","pdp_static.common.my_orders":"รายการสั่งซื้อของฉัน","pdp_static.common.noDesc":"ไม่มีรายละเอียด","pdp_static.common.no_desc":"ไม่มีรายละเอียด","pdp_static.common.quantity":"จำนวน","pdp_static.common.register":"สมัครสมาชิก","pdp_static.common.requestFailureDefaultMessage":"เกิดข้อผิดพลาดในระบบ กรุณาลองใหม่อีกครั้งภายหลัง","pdp_static.common.request_failure_default_message":"เกิดข้อผิดพลาดในระบบ กรุณาลองใหม่อีกครั้งภายหลัง","pdp_static.common.share_via":"แชร์โดย","pdp_static.common.star":"ดาว","pdp_static.common.tryagain":"ลองใหม่อีกครั้ง","pdp_static.common.video_error":"เกิดข้อผิดพลาดในระบบ กรุณาลองใหม่อีกครั้งภายหลัง","pdp_static.common.viewAll":"ดูทั้งหมด","pdp_static.common.viewLess":"ย่อรายละเอียด","pdp_static.common.viewMore":"ดูเพิ่มเติม","pdp_static.common.view_all":"ดูทั้งหมด","pdp_static.common.view_less":"ย่อรายละเอียด","pdp_static.common.view_more":"ดูเพิ่ม","pdp_static.delivery.change":"เปลี่ยน","pdp_static.delivery.free":"ฟรี","pdp_static.delivery.getBy":"ได้รับภายใน {time}","pdp_static.delivery.get_by":"ได้รับภายใน {time}","pdp_static.delivery.location_notice":"ไม่สามารถส่งรายการสินค้านี้ไปยังสถานที่นี้ กรุณาเลือกสถานที่อื่น","pdp_static.delivery.save":"บันทึก","pdp_static.delivery.title":"ตัวเลือกการจัดส่ง","pdp_static.description.title":"Description","pdp_static.dialog.gotIt":"ยืนยัน","pdp_static.dialog.got_it":"ยืนยัน","pdp_static.disclaimer.title":"Disclaimer","pdp_static.emi_popup.content":"Test","pdp_static.error.tip.item.notfound":"เกิดข้อผิดพลาด \\nไม่พบสินค้าที่ท่านต้องการ","pdp_static.error.tip.try.again":"เกิดข้อผิดพลาด \\nกรุณาตรวจสอบการเชื่อมต่อของท่านและลองอีกครั้ง","pdp_static.freeGift.chooseTitle":"เลือกของขวัญฟรี","pdp_static.freeGift.free":"ฟรี","pdp_static.freeGift.title":"ฟรีของขวัญ","pdp_static.freeSample.chooseTitle":"กรุณาเลือกสินค้าตัวอย่างฟรี","pdp_static.freeSample.title":"ฟรีสินค้าตัวอย่าง","pdp_static.free_gift.choose_title":"เลือกของขวัญฟรี","pdp_static.free_gift.free":"ฟรี","pdp_static.free_gift.title":"ฟรีของขวัญ","pdp_static.free_sample.choose_title":"กรุณาเลือกสินค้าตัวอย่างฟรี","pdp_static.free_sample.title":"ฟรีสินค้าตัวอย่าง","pdp_static.go_to_wishlist":"ไปยังรายการที่ชอบ","pdp_static.groupbuy.complete":"Groupbuy สำเร็จ","pdp_static.groupbuy.create":"สร้างกลุ่ม","pdp_static.groupbuy.endin":"สิ้นสุดใน","pdp_static.groupbuy.ends":"ends on","pdp_static.groupbuy.for":"จำนวน","pdp_static.groupbuy.invite":"เชิญเพื่อน","pdp_static.groupbuy.join":"เข้าร่วมกลุ่ม","pdp_static.groupbuy.listtitle":"เข้าร่วมกลุ่ม","pdp_static.groupbuy.okBtnText":"ตกลง","pdp_static.groupbuy.ruletext1":"คำสั่งซื้อจะยังไม่ดำเนินการ จนกว่าจะสร้างกลุ่มชวนเพื่อนร่วมกันซื้อสำเร็จ ก่อนแคมเปญจบ","pdp_static.groupbuy.ruletext2":"จะดำเนินการคืนเงินให้ ถ้ารวมกลุ่มชวนเพื่อนร่วมกันซื้อไม่สำเร็จ","pdp_static.groupbuy.ruletext3":"กลุ่มชวนเพื่อนร่วมกันซื้อไม่สามารถ ชำระเงินปลายทาง หรือ ชำระผ่านเคาน์เตอร์จ่ายเงินได้","pdp_static.groupbuy.ruletitle":"กฏในการเข้าร่วมกลุ่มชวนเพื่อนร่วมกันซื้อ","pdp_static.groupbuy.ship":"สินค้าส่งแล้ว","pdp_static.groupbuy.viewrule":"อ่านกฏในการเข่าร่วมกลุ่มชวนเพื่อนร่วมกันซื้อ","pdp_static.helpful_statement_like":"{likeCount, plural, =0 {You found this helpful} one {You and one person found this helpful} other {You and {likeCount} people found this helpful}}","pdp_static.highlights.title":"Highlights","pdp_static.installment.text":"สูงสุดถึง {months} เดือน ผ่อนงวดละ {amount} ต่อเดือน","pdp_static.installment.title":"การผ่อนชำระ","pdp_static.ms.countdown_mega":"จะสิ้นสุดภายใน {time}","pdp_static.ms.countdown_mega_day":"จะสิ้นสุดภายใน {days} วัน {time}","pdp_static.ms.countdown_mega_days":"จะสิ้นสุดภายใน {days} วัน {time}\n","pdp_static.ms.countdown_teaser":"จะเริ่มภายใน {days} วัน {time}","pdp_static.ms.countdown_teaser_no_day":"จะเริ่มภายใน  {time}","pdp_static.ms.freeShipping":"ฟรีค่าขนส่ง","pdp_static.ms.…