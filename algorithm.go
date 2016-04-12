package ptlogin

var algorithm map[string]string = make(map[string]string)

func init() {
	algorithm["TEA"] = `
    function r(t) {
        for (var e = "", n = 0; n < t.length; n += 2)e += String.fromCharCode(parseInt(t.substr(n, 2), 16));
        return e
    }
    function a(t, e) {
        if (!t)return "";
        e && (t = p(t));
        for (var n = [], i = 0; i < t.length; i++)n[i] = t.charCodeAt(i);
        return o(n)
    }
    function p(t) {
        var e, n, i = [], o = t.length;
        for (e = 0; o > e; e++)n = t.charCodeAt(e), n > 0 && 127 >= n ? i.push(t.charAt(e)) : n >= 128 && 2047 >= n ? i.push(String.fromCharCode(192 | n >> 6 & 31), String.fromCharCode(128 | 63 & n)) : n >= 2048 && 65535 >= n && i.push(String.fromCharCode(224 | n >> 12 & 15), String.fromCharCode(128 | n >> 6 & 63), String.fromCharCode(128 | 63 & n));
        return i.join("")
    }
    function f(t, e) {
        var n = [];
        if (e)for (var i = 0; i < t.length; i++)n[i] = 255 & t.charCodeAt(i); else for (var o = 0, i = 0; i < t.length; i += 2)n[o++] = parseInt(t.substr(i, 2), 16);
        return n
    }
    function o(t) {
        if (!t)return "";
        for (var e = "", n = 0; n < t.length; n++) {
            var i = Number(t[n]).toString(16);
            1 == i.length && (i = "0" + i), e += i
        }
        return e
    }

    function e() {
        //return 4110171141
        return Math.round(4294967295 * Math.random())
    }

    function l(t) {
        for (var e = 16, o = n(t, 0, 4), r = n(t, 4, 4), a = n(g, 0, 4), p = n(g, 4, 4), s = n(g, 8, 4), c = n(g, 12, 4), u = 0, l = 2654435769; e-- > 0;)u += l, u = (4294967295 & u) >>> 0, o += (r << 4) + a ^ r + u ^ (r >>> 5) + p, o = (4294967295 & o) >>> 0, r += (o << 4) + s ^ o + u ^ (o >>> 5) + c, r = (4294967295 & r) >>> 0;
        var d = new Array(8);
        return i(d, 0, o), i(d, 4, r), d
    }
    function n(t, e, n) {
        (!n || n > 4) && (n = 4);
        for (var i = 0, o = e; e + n > o; o++)i <<= 8, i |= t[o];
        return (4294967295 & i) >>> 0
    }

    function i(t, e, n) {
        t[e + 3] = n >> 0 & 255, t[e + 2] = n >> 8 & 255, t[e + 1] = n >> 16 & 255, t[e + 0] = n >> 24 & 255
    }

    function u() {
        for (var t = 0; 8 > t; t++)_[t] ^= $ ? w[t] : b[y + t];
        for (var e = l(_), t = 0; 8 > t; t++)b[v + t] = e[t] ^ w[t], w[t] = _[t];
        y = v, v += 8, m = 0, $ = !1
    }



    function s(t) {
        _ = new Array(8), w = new Array(8), v = y = 0, $ = !0, m = 0;
        var n = t.length, i = 0;
        m = (n + 10) % 8, 0 != m && (m = 8 - m), b = new Array(n + m + 10), _[0] = 255 & (248 & e() | m);
        for (var o = 1; m >= o; o++)_[o] = 255 & e();
        m++;
        for (var o = 0; 8 > o; o++)w[o] = 0;
        for (i = 1; 2 >= i;)8 > m && (_[m++] = 255 & e(), i++), 8 == m && u();
        for (var o = 0; n > 0;)8 > m && (_[m++] = t[o++], n--), 8 == m && u();
        for (i = 1; 7 >= i;)8 > m && (_[m++] = 0, i++), 8 == m && u();
        return b
    }


    var g = "", m = 0, _ = [], w = [], v = 0, y = 0, b = [], k = [], $ = !0;
    TEA = {
        encrypt: function (t, e) {
            var n = f(t, e), i = s(n);
            return o(i)
        }, enAsBase64: function (t, e) {
            for (var n = f(t, e), i = s(n), o = "", r = 0; r < i.length; r++)o += String.fromCharCode(i[r]);
            return btoa(o)
        }, decrypt: function (t) {
            var e = f(t, !1), n = c(e);
            return o(n)
        }, initkey: function (t, e) {
            g = f(t, e)
        }, bytesToStr: r, strToBytes: a, bytesInStr: o, dataFromStr: f
    };

    //console.log("js key:",key)
    TEA.initkey(key);
    //console.log("js text:",plaint_text)
    result = TEA.encrypt(plaint_text);
    //console.log("js result:",result)
`
	algorithm["RSA"] = `
RSA = function () {
    function t(t, e) {
        return new a(t, e)
    }

    function e(t, e) {
        if (e < t.length + 11)return uv_alert("Message too long for RSA"), null;
        for (var n = new Array, i = t.length - 1; i >= 0 && e > 0;) {
            var o = t.charCodeAt(i--);
            n[--e] = o
        }
        n[--e] = 0;
        for (var r = new Y, p = new Array; e > 2;) {
            for (p[0] = 0; 0 == p[0];)r.nextBytes(p);
            n[--e] = p[0]
        }
        return n[--e] = 2, n[--e] = 0, new a(n)
    }

    function n() {
        this.n = null, this.e = 0, this.d = null, this.p = null, this.q = null, this.dmp1 = null, this.dmq1 = null, this.coeff = null
    }

    function i(e, n) {
        null != e && null != n && e.length > 0 && n.length > 0 ? (this.n = t(e, 16), this.e = parseInt(n, 16)) : uv_alert("Invalid RSA public key")
    }

    function o(t) {
        return t.modPowInt(this.e, this.n)
    }

    function r(t) {
        var n = e(t, this.n.bitLength() + 7 >> 3);
        if (null == n)return null;
        var i = this.doPublic(n);
        if (null == i)return null;
        var o = i.toString(16);
        return 0 == (1 & o.length) ? o : "0" + o
    }

    function a(t, e, n) {
        null != t && ("number" == typeof t ? this.fromNumber(t, e, n) : null == e && "string" != typeof t ? this.fromString(t, 256) : this.fromString(t, e))
    }

    function p() {
        return new a(null)
    }

    function s(t, e, n, i, o, r) {
        for (; --r >= 0;) {
            var a = e * this[t++] + n[i] + o;
            o = Math.floor(a / 67108864), n[i++] = 67108863 & a
        }
        return o
    }

    function c(t, e, n, i, o, r) {
        for (var a = 32767 & e, p = e >> 15; --r >= 0;) {
            var s = 32767 & this[t], c = this[t++] >> 15, u = p * s + c * a;
            s = a * s + ((32767 & u) << 15) + n[i] + (1073741823 & o), o = (s >>> 30) + (u >>> 15) + p * c + (o >>> 30), n[i++] = 1073741823 & s
        }
        return o
    }

    function u(t, e, n, i, o, r) {
        for (var a = 16383 & e, p = e >> 14; --r >= 0;) {
            var s = 16383 & this[t], c = this[t++] >> 14, u = p * s + c * a;
            s = a * s + ((16383 & u) << 14) + n[i] + o, o = (s >> 28) + (u >> 14) + p * c, n[i++] = 268435455 & s
        }
        return o
    }

    function l(t) {
        return lt.charAt(t)
    }

    function d(t, e) {
        var n = dt[t.charCodeAt(e)];
        return null == n ? -1 : n
    }

    function h(t) {
        for (var e = this.t - 1; e >= 0; --e)t[e] = this[e];
        t.t = this.t, t.s = this.s
    }

    function f(t) {
        this.t = 1, this.s = 0 > t ? -1 : 0, t > 0 ? this[0] = t : -1 > t ? this[0] = t + DV : this.t = 0
    }

    function g(t) {
        var e = p();
        return e.fromInt(t), e
    }

    function m(t, e) {
        var n;
        if (16 == e)n = 4; else if (8 == e)n = 3; else if (256 == e)n = 8; else if (2 == e)n = 1; else if (32 == e)n = 5; else {
            if (4 != e)return void this.fromRadix(t, e);
            n = 2
        }
        this.t = 0, this.s = 0;
        for (var i = t.length, o = !1, r = 0; --i >= 0;) {
            var p = 8 == n ? 255 & t[i] : d(t, i);
            0 > p ? "-" == t.charAt(i) && (o = !0) : (o = !1, 0 == r ? this[this.t++] = p : r + n > this.DB ? (this[this.t - 1] |= (p & (1 << this.DB - r) - 1) << r, this[this.t++] = p >> this.DB - r) : this[this.t - 1] |= p << r, r += n, r >= this.DB && (r -= this.DB))
        }
        8 == n && 0 != (128 & t[0]) && (this.s = -1, r > 0 && (this[this.t - 1] |= (1 << this.DB - r) - 1 << r)), this.clamp(), o && a.ZERO.subTo(this, this)
    }

    function _() {
        for (var t = this.s & this.DM; this.t > 0 && this[this.t - 1] == t;)--this.t
    }

    function w(t) {
        if (this.s < 0)return "-" + this.negate().toString(t);
        var e;
        if (16 == t)e = 4; else if (8 == t)e = 3; else if (2 == t)e = 1; else if (32 == t)e = 5; else {
            if (4 != t)return this.toRadix(t);
            e = 2
        }
        var n, i = (1 << e) - 1, o = !1, r = "", a = this.t, p = this.DB - a * this.DB % e;
        if (a-- > 0)for (p < this.DB && (n = this[a] >> p) > 0 && (o = !0, r = l(n)); a >= 0;)e > p ? (n = (this[a] & (1 << p) - 1) << e - p, n |= this[--a] >> (p += this.DB - e)) : (n = this[a] >> (p -= e) & i, 0 >= p && (p += this.DB, --a)), n > 0 && (o = !0), o && (r += l(n));
        return o ? r : "0"
    }

    function v() {
        var t = p();
        return a.ZERO.subTo(this, t), t
    }

    function y() {
        return this.s < 0 ? this.negate() : this
    }

    function b(t) {
        var e = this.s - t.s;
        if (0 != e)return e;
        var n = this.t;
        if (e = n - t.t, 0 != e)return e;
        for (; --n >= 0;)if (0 != (e = this[n] - t[n]))return e;
        return 0
    }

    function k(t) {
        var e, n = 1;
        return 0 != (e = t >>> 16) && (t = e, n += 16), 0 != (e = t >> 8) && (t = e, n += 8), 0 != (e = t >> 4) && (t = e, n += 4), 0 != (e = t >> 2) && (t = e, n += 2), 0 != (e = t >> 1) && (t = e, n += 1), n
    }

    function $() {
        return this.t <= 0 ? 0 : this.DB * (this.t - 1) + k(this[this.t - 1] ^ this.s & this.DM)
    }

    function q(t, e) {
        var n;
        for (n = this.t - 1; n >= 0; --n)e[n + t] = this[n];
        for (n = t - 1; n >= 0; --n)e[n] = 0;
        e.t = this.t + t, e.s = this.s
    }

    function S(t, e) {
        for (var n = t; n < this.t; ++n)e[n - t] = this[n];
        e.t = Math.max(this.t - t, 0), e.s = this.s
    }

    function A(t, e) {
        var n, i = t % this.DB, o = this.DB - i, r = (1 << o) - 1, a = Math.floor(t / this.DB), p = this.s << i & this.DM;
        for (n = this.t - 1; n >= 0; --n)e[n + a + 1] = this[n] >> o | p, p = (this[n] & r) << i;
        for (n = a - 1; n >= 0; --n)e[n] = 0;
        e[a] = p, e.t = this.t + a + 1, e.s = this.s, e.clamp()
    }

    function T(t, e) {
        e.s = this.s;
        var n = Math.floor(t / this.DB);
        if (n >= this.t)return void(e.t = 0);
        var i = t % this.DB, o = this.DB - i, r = (1 << i) - 1;
        e[0] = this[n] >> i;
        for (var a = n + 1; a < this.t; ++a)e[a - n - 1] |= (this[a] & r) << o, e[a - n] = this[a] >> i;
        i > 0 && (e[this.t - n - 1] |= (this.s & r) << o), e.t = this.t - n, e.clamp()
    }

    function I(t, e) {
        for (var n = 0, i = 0, o = Math.min(t.t, this.t); o > n;)i += this[n] - t[n], e[n++] = i & this.DM, i >>= this.DB;
        if (t.t < this.t) {
            for (i -= t.s; n < this.t;)i += this[n], e[n++] = i & this.DM, i >>= this.DB;
            i += this.s
        } else {
            for (i += this.s; n < t.t;)i -= t[n], e[n++] = i & this.DM, i >>= this.DB;
            i -= t.s
        }
        e.s = 0 > i ? -1 : 0, -1 > i ? e[n++] = this.DV + i : i > 0 && (e[n++] = i), e.t = n, e.clamp()
    }

    function E(t, e) {
        var n = this.abs(), i = t.abs(), o = n.t;
        for (e.t = o + i.t; --o >= 0;)e[o] = 0;
        for (o = 0; o < i.t; ++o)e[o + n.t] = n.am(0, i[o], e, o, 0, n.t);
        e.s = 0, e.clamp(), this.s != t.s && a.ZERO.subTo(e, e)
    }

    function C(t) {
        for (var e = this.abs(), n = t.t = 2 * e.t; --n >= 0;)t[n] = 0;
        for (n = 0; n < e.t - 1; ++n) {
            var i = e.am(n, e[n], t, 2 * n, 0, 1);
            (t[n + e.t] += e.am(n + 1, 2 * e[n], t, 2 * n + 1, i, e.t - n - 1)) >= e.DV && (t[n + e.t] -= e.DV, t[n + e.t + 1] = 1)
        }
        t.t > 0 && (t[t.t - 1] += e.am(n, e[n], t, 2 * n, 0, 1)), t.s = 0, t.clamp()
    }

    function D(t, e, n) {
        var i = t.abs();
        if (!(i.t <= 0)) {
            var o = this.abs();
            if (o.t < i.t)return null != e && e.fromInt(0), void(null != n && this.copyTo(n));
            null == n && (n = p());
            var r = p(), s = this.s, c = t.s, u = this.DB - k(i[i.t - 1]);
            u > 0 ? (i.lShiftTo(u, r), o.lShiftTo(u, n)) : (i.copyTo(r), o.copyTo(n));
            var l = r.t, d = r[l - 1];
            if (0 != d) {
                var h = d * (1 << this.F1) + (l > 1 ? r[l - 2] >> this.F2 : 0), f = this.FV / h, g = (1 << this.F1) / h, m = 1 << this.F2, _ = n.t, w = _ - l, v = null == e ? p() : e;
                for (r.dlShiftTo(w, v), n.compareTo(v) >= 0 && (n[n.t++] = 1, n.subTo(v, n)), a.ONE.dlShiftTo(l, v), v.subTo(r, r); r.t < l;)r[r.t++] = 0;
                for (; --w >= 0;) {
                    var y = n[--_] == d ? this.DM : Math.floor(n[_] * f + (n[_ - 1] + m) * g);
                    if ((n[_] += r.am(0, y, n, w, 0, l)) < y)for (r.dlShiftTo(w, v), n.subTo(v, n); n[_] < --y;)n.subTo(v, n)
                }
                null != e && (n.drShiftTo(l, e), s != c && a.ZERO.subTo(e, e)), n.t = l, n.clamp(), u > 0 && n.rShiftTo(u, n), 0 > s && a.ZERO.subTo(n, n)
            }
        }
    }

    function B(t) {
        var e = p();
        return this.abs().divRemTo(t, null, e), this.s < 0 && e.compareTo(a.ZERO) > 0 && t.subTo(e, e), e
    }

    function M(t) {
        this.m = t
    }

    function x(t) {
        return t.s < 0 || t.compareTo(this.m) >= 0 ? t.mod(this.m) : t
    }

    function L(t) {
        return t
    }

    function O(t) {
        t.divRemTo(this.m, null, t)
    }

    function H(t, e, n) {
        t.multiplyTo(e, n), this.reduce(n)
    }

    function R(t, e) {
        t.squareTo(e), this.reduce(e)
    }

    function K() {
        if (this.t < 1)return 0;
        var t = this[0];
        if (0 == (1 & t))return 0;
        var e = 3 & t;
        return e = e * (2 - (15 & t) * e) & 15, e = e * (2 - (255 & t) * e) & 255, e = e * (2 - ((65535 & t) * e & 65535)) & 65535, e = e * (2 - t * e % this.DV) % this.DV, e > 0 ? this.DV - e : -e
    }

    function U(t) {
        this.m = t, this.mp = t.invDigit(), this.mpl = 32767 & this.mp, this.mph = this.mp >> 15, this.um = (1 << t.DB - 15) - 1, this.mt2 = 2 * t.t
    }

    function N(t) {
        var e = p();
        return t.abs().dlShiftTo(this.m.t, e), e.divRemTo(this.m, null, e), t.s < 0 && e.compareTo(a.ZERO) > 0 && this.m.subTo(e, e), e
    }

    function P(t) {
        var e = p();
        return t.copyTo(e), this.reduce(e), e
    }

    function F(t) {
        for (; t.t <= this.mt2;)t[t.t++] = 0;
        for (var e = 0; e < this.m.t; ++e) {
            var n = 32767 & t[e], i = n * this.mpl + ((n * this.mph + (t[e] >> 15) * this.mpl & this.um) << 15) & t.DM;
            for (n = e + this.m.t, t[n] += this.m.am(0, i, t, e, 0, this.m.t); t[n] >= t.DV;)t[n] -= t.DV, t[++n]++
        }
        t.clamp(), t.drShiftTo(this.m.t, t), t.compareTo(this.m) >= 0 && t.subTo(this.m, t)
    }

    function Q(t, e) {
        t.squareTo(e), this.reduce(e)
    }

    function j(t, e, n) {
        t.multiplyTo(e, n), this.reduce(n)
    }

    function V() {
        return 0 == (this.t > 0 ? 1 & this[0] : this.s)
    }

    function z(t, e) {
        if (t > 4294967295 || 1 > t)return a.ONE;
        var n = p(), i = p(), o = e.convert(this), r = k(t) - 1;
        for (o.copyTo(n); --r >= 0;)if (e.sqrTo(n, i), (t & 1 << r) > 0)e.mulTo(i, o, n); else {
            var s = n;
            n = i, i = s
        }
        return e.revert(n)
    }

    function G(t, e) {
        var n;
        return n = 256 > t || e.isEven() ? new M(e) : new U(e), this.exp(t, n)
    }

    function J(t) {
        ft[gt++] ^= 255 & t, ft[gt++] ^= t >> 8 & 255, ft[gt++] ^= t >> 16 & 255, ft[gt++] ^= t >> 24 & 255, gt >= wt && (gt -= wt)
    }

    function W() {
        //J(1460365486008)
        J((new Date).getTime())
    }

    function Z() {
        if (null == ht) {
            for (W(), ht = it(), ht.init(ft), gt = 0; gt < ft.length; ++gt)ft[gt] = 0;
            gt = 0
        }
        return ht.next()
    }

    function X(t) {
        var e;
        for (e = 0; e < t.length; ++e)t[e] = Z()
    }

    function Y() {
    }

    function tt() {
        this.i = 0, this.j = 0, this.S = new Array
    }

    function et(t) {
        var e, n, i;
        for (e = 0; 256 > e; ++e)this.S[e] = e;
        for (n = 0, e = 0; 256 > e; ++e)n = n + this.S[e] + t[e % t.length] & 255, i = this.S[e], this.S[e] = this.S[n], this.S[n] = i;
        this.i = 0, this.j = 0
    }

    function nt() {
        var t;
        return this.i = this.i + 1 & 255, this.j = this.j + this.S[this.i] & 255, t = this.S[this.i], this.S[this.i] = this.S[this.j], this.S[this.j] = t, this.S[t + this.S[this.i] & 255]
    }

    function it() {
        return new tt
    }

    function ot(t, e, i) {
        e = "F20CE00BAE5361F8FA3AE9CEFA495362FF7DA1BA628F64A347F0A8C012BF0B254A30CD92ABFFE7A6EE0DC424CB6166F8819EFA5BCCB20EDFB4AD02E412CCF579B1CA711D55B8B0B3AEB60153D5E0693A2A86F3167D7847A0CB8B00004716A9095D9BADC977CBB804DBDCBA6029A9710869A453F27DFDDF83C016D928B3CBF4C7", i = "3";
        var o = new n;
        return o.setPublic(e, i), o.encrypt(t)
    }

    n.prototype.doPublic = o, n.prototype.setPublic = i, n.prototype.encrypt = r;
    var rt, at = 0xdeadbeefcafe, pt = 15715070 == (16777215 & at);
    (a.prototype.am = u, rt = 28), a.prototype.DB = rt, a.prototype.DM = (1 << rt) - 1, a.prototype.DV = 1 << rt;
    var st = 52;
    a.prototype.FV = Math.pow(2, st), a.prototype.F1 = st - rt, a.prototype.F2 = 2 * rt - st;
    var ct, ut, lt = "0123456789abcdefghijklmnopqrstuvwxyz", dt = new Array;
    for (ct = "0".charCodeAt(0), ut = 0; 9 >= ut; ++ut)dt[ct++] = ut;
    for (ct = "a".charCodeAt(0), ut = 10; 36 > ut; ++ut)dt[ct++] = ut;
    for (ct = "A".charCodeAt(0), ut = 10; 36 > ut; ++ut)dt[ct++] = ut;
    M.prototype.convert = x, M.prototype.revert = L, M.prototype.reduce = O, M.prototype.mulTo = H, M.prototype.sqrTo = R, U.prototype.convert = N, U.prototype.revert = P, U.prototype.reduce = F, U.prototype.mulTo = j, U.prototype.sqrTo = Q, a.prototype.copyTo = h, a.prototype.fromInt = f, a.prototype.fromString = m, a.prototype.clamp = _, a.prototype.dlShiftTo = q, a.prototype.drShiftTo = S, a.prototype.lShiftTo = A, a.prototype.rShiftTo = T, a.prototype.subTo = I, a.prototype.multiplyTo = E, a.prototype.squareTo = C, a.prototype.divRemTo = D, a.prototype.invDigit = K, a.prototype.isEven = V, a.prototype.exp = z, a.prototype.toString = w, a.prototype.negate = v, a.prototype.abs = y, a.prototype.compareTo = b, a.prototype.bitLength = $, a.prototype.mod = B, a.prototype.modPowInt = G, a.ZERO = g(0), a.ONE = g(1);
    var ht, ft, gt;
    if (null == ft) {
        ft = new Array, gt = 0;
        var mt;
        for (; wt > gt;)mt = 29853, ft[gt++] = mt >>> 8, ft[gt++] = 255 & mt;
        gt = 0, W()
    }
    Y.prototype.nextBytes = X, tt.prototype.init = et, tt.prototype.next = nt;
    var wt = 256;
    return {rsa_encrypt: ot}
}()

function hexchar2bin(str) {
    for (var arr = [], i = 0; i < str.length; i += 2)arr.push("\\x" + str.substr(i, 2));
    return arr = arr.join(""), eval("var temp = '" + arr + "'"), temp
}
    r = hexchar2bin(plaint_text)
    p = RSA.rsa_encrypt(r)
    `

}
