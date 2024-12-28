/**
 * puzzles/2024/day23/solution.ts
 *
 * ~~ LAN Party ~~
 * this is my solution for this advent of code puzzle
 *
 * by alex prosser
 * 12/22/2024
 */

// find all combinations of array with k entries
const combination = (array: string[], k: number) => {
  const result: string[][] = [];

  const helper = (
    _array: string[],
    _k: number,
    _i: number,
    _current: string[]
  ) => {
    if (_current.length == k) result.push(_current);
    if (_current.length == k || _i == _array.length) return;

    helper(_array, _k, _i + 1, [_array[_i], ..._current]);
    helper(_array, _k, _i + 1, [..._current]);
  };

  helper(array, k, 0, []);
  return result;
};

// recursively find the set that can loop given a size time
// note: this function only works for size 3
const findTripleSet = (
  graph: { [key: string]: string[] },
  path: string[]
): string[][] => {
  const current = path.at(-1) as string;
  if (path.length === 4) {
    if (current === path[0]) return [path.slice(0, 3)];
    else return [];
  }

  // check for repeats
  if (new Set(path).size !== path.length) return [];

  const allSets: string[][] = [];
  for (let i = 0; i < graph[current].length; i++) {
    path.push(graph[current][i]);
    const sets = findTripleSet(graph, path);
    path.pop();

    if (sets.length > 0) allSets.push(...sets);
  }
  return allSets;
};

/**
 * the code of part 1 of the puzzle
 */
const part1 = (input: string) => {
  const graph = input
    .trim()
    .split("\n")
    .reduce<{ [key: string]: string[] }>((obj, line) => {
      const [left, right] = line.split("-");
      if (obj[left] === undefined) obj[left] = [];
      if (obj[right] === undefined) obj[right] = [];

      obj[left].push(right);
      obj[right].push(left);
      return obj;
    }, {});

  let allSets = new Set<string>();
  Object.keys(graph).forEach((node) => {
    allSets = allSets.union(
      new Set(findTripleSet(graph, [node]).map((set) => set.sort().join(",")))
    );
  });

  console.log(
    Array.from(allSets).reduce((sum, set) => {
      if (set.split(",").find((node) => node.startsWith("t")) !== undefined)
        sum++;
      return sum;
    }, 0)
  );
  return Array.from(allSets).reduce((sum, set) => {
    if (set.split(",").find((node) => node.startsWith("t")) !== undefined)
      sum++;
    return sum;
  }, 0);
};

part1(`rv-jv
gb-fa
uc-lw
ij-tw
vh-ab
rk-dp
iv-oi
uu-kl
rs-zb
gk-gb
bl-gk
ea-au
gi-bt
ef-qi
ey-hv
wi-of
yd-sv
gn-fc
eb-af
ru-tg
ds-by
hw-pd
ma-on
en-kl
cw-do
vj-ou
ab-ae
oe-ca
zv-ru
fp-io
xa-wg
vm-gg
sb-fo
vm-ot
up-jc
ui-jj
ww-im
op-tk
nh-yg
vc-ou
pr-jz
oy-zh
pn-hr
xy-ud
ip-jo
wo-xl
xy-ab
bi-bl
vv-gu
mf-pi
pi-nn
gc-kv
zi-hm
bs-yd
qs-jv
tg-lm
ag-ug
kz-fd
bh-hy
qg-vv
ul-cj
xi-of
vy-jp
py-va
bb-pf
yb-oo
nd-cy
yi-tm
ou-bt
nf-pc
fc-fp
tu-xy
qk-hp
ac-ua
un-lh
vc-vj
vg-fy
hy-ko
jt-pt
vi-mt
td-fc
iv-fu
fn-fy
dm-jb
fu-wz
br-ql
rs-ey
mr-gh
ya-rv
an-jc
ze-tl
io-fz
dz-ce
dg-au
kq-yx
yn-eg
ar-hs
ua-yx
fd-wp
xx-bc
rn-wo
zj-ao
zy-ts
ld-wz
wi-pl
ta-nk
if-bm
ia-wt
df-jf
wq-zf
lz-qq
mf-ts
wz-oi
qv-oa
pr-xz
ih-ya
vb-ks
hg-yw
zc-cw
ir-di
wh-xp
hw-lp
km-cj
mt-hx
ja-hd
ul-nx
ef-yw
cs-cc
dk-wy
uz-ar
ai-pg
wb-ci
gr-ll
hq-ox
lj-kl
ee-xs
sj-oo
gz-ky
vx-wp
xp-vm
lb-lv
qe-rj
yu-rm
fy-bo
je-wo
hv-uk
cc-fu
ut-sq
zw-fx
ai-kt
ya-jz
ob-jv
ix-zz
rx-jo
qc-in
cz-bh
op-rf
hu-yd
tu-rt
uo-yj
el-xy
xn-oy
yv-gg
qd-sl
oo-gk
ld-zc
lr-ng
en-sw
pi-nt
uz-xb
xd-wa
cx-gq
gd-ns
qv-nh
jo-sz
pe-zo
sz-ip
zk-xs
sp-bd
bb-mv
vm-wk
re-ho
ei-ae
av-bo
zr-cl
ib-kz
ez-ob
cq-dd
fn-xh
lh-su
pk-zv
gj-vp
hz-my
re-tb
st-zy
qf-no
wz-az
ls-nt
hx-lw
ac-lw
qn-rq
of-pw
xy-ei
qq-ao
mq-gq
cw-az
cc-wz
gh-en
kp-dx
aa-bf
wm-yr
hx-ir
la-vz
tw-ou
zm-qq
gu-kf
in-kz
eh-oq
lj-xz
vp-it
zs-hu
ts-er
ce-up
ma-st
nl-gj
nl-vp
zw-nq
yn-hm
rn-aw
sz-fm
wg-hn
st-sy
qn-pc
zg-gn
tn-jj
zh-xn
hu-wq
fu-cw
eq-py
po-sg
yt-cj
dl-un
nb-dv
dw-sb
dp-rt
uq-ja
kq-bg
pv-nc
sg-lg
tc-rf
zd-nm
mx-lg
xh-fy
jw-vp
rb-ow
fi-xn
os-qc
pk-ng
kr-id
co-fh
dj-vo
ps-kd
qh-wx
wf-nb
li-up
or-wx
jl-cq
xv-hd
yo-rm
wg-su
hu-vt
kn-ol
qp-qx
aw-je
nv-np
ob-tw
pw-rt
zj-bv
yd-zs
fo-xr
qp-tq
xl-ul
yb-gk
bh-sa
yw-mt
ns-rl
kp-is
xe-ef
jc-lx
yu-nd
nb-nk
ji-oe
ul-pn
we-rv
yt-pn
tx-jm
wc-qk
ql-zw
xz-hg
nh-ik
sg-ae
ey-ro
pg-ne
tz-oi
sl-wq
yy-ng
ol-gy
la-uc
xl-nx
xl-rn
tu-ad
ih-jv
nt-dd
ug-kk
zw-jt
kx-jc
tz-uv
sv-zs
mh-wn
rv-ih
dz-lb
gx-ey
lt-zc
he-lh
lo-ou
vz-ko
eu-qi
ft-sf
wr-sl
xh-sn
ik-kh
tq-df
hw-ac
kr-yf
dm-qr
eh-jp
hk-wt
vy-ml
ix-va
kx-lx
ky-fs
di-et
bm-kd
ha-ou
ir-lu
db-ba
iy-qn
ir-et
io-zg
uo-eg
mf-on
hu-rq
py-sm
vi-ff
tx-hz
or-ul
xe-co
pw-nn
jg-ap
ee-dx
dw-ji
mf-gc
gb-ba
yp-lv
zm-cb
xy-ae
kz-bc
yd-rq
gq-hp
yx-mq
ag-hs
rz-kr
dz-an
nj-lu
ng-ms
nt-jl
ft-ui
ja-pj
ef-co
ea-ta
uw-tl
zf-sl
it-nn
an-uu
cj-jx
in-fd
cl-hp
yc-hd
wo-wk
hk-tz
op-ee
cb-do
aj-tp
cv-sf
va-zz
mx-xy
wz-om
xh-yi
uw-rj
bo-fn
hf-ni
gd-vx
bd-ln
pw-tm
ar-mj
qq-hk
vp-bm
rq-zf
ms-gu
nq-qr
ra-ko
pq-ap
wo-cj
lp-pz
vo-dv
tj-ab
uw-ml
fm-ip
wk-wh
rg-ip
wq-bs
dh-tp
aw-km
wt-zm
xd-sz
ts-gc
gq-re
pz-ub
pj-ju
oq-os
wl-cb
xm-ws
rm-hm
ag-uz
ps-gy
fg-yl
cy-wh
bz-ln
mk-oy
vo-wj
oq-ik
qx-fn
uq-pj
cb-uv
ro-ya
lu-vb
om-ld
lz-qp
kx-up
ir-nj
kh-cc
nf-bq
ft-ds
vx-xx
uq-ju
cw-fi
jw-ol
rt-nn
wz-ex
zw-pt
xl-ib
ad-qw
dj-cx
az-oi
ks-ji
qn-tx
wj-gl
cw-ld
uk-or
sw-mr
yg-qv
xk-mu
qo-kd
wi-jh
ti-yw
dz-wl
vv-yy
xm-fg
ci-er
ox-kk
gq-zr
pr-rv
bm-gy
nq-jt
zp-uw
lx-dz
fz-ws
wo-jx
lu-hz
xo-pq
ly-cq
gh-pg
xs-kp
fp-fz
fm-uy
mr-bw
wo-ul
cc-az
gx-yo
rl-kz
bx-iu
yw-vi
eu-ti
vr-fo
kk-ag
lp-kq
ca-ks
az-fu
zb-hv
ti-ff
jh-sp
dq-wr
wz-cs
nk-vo
nx-yt
tv-cx
mj-kk
no-ia
ws-fg
gc-er
jo-fm
ee-tc
ds-sf
eb-yr
on-gc
xp-mu
wa-ut
oq-vy
iv-om
xo-eq
ww-ly
hf-fg
hh-pk
qi-td
am-lo
dq-qd
lm-yf
bn-xz
yv-uc
ow-rs
zi-uo
rb-gx
tv-dj
sj-ap
oj-gy
vj-ij
sb-if
yp-np
of-zq
dh-jg
th-yc
yk-kk
tt-sy
he-iu
ko-cr
lp-ry
cn-uc
rf-pv
py-ap
gc-sy
ng-bp
dg-ta
tk-pv
gr-xn
ns-nc
ca-kg
ni-em
rq-zs
wv-ay
uj-tn
yx-lp
jw-ps
dv-gl
vx-in
yr-xp
fm-pc
nq-mu
yj-op
wh-eb
zc-cs
cr-jj
kn-it
sm-vt
oo-fa
ud-oe
ei-tu
xe-hx
is-tk
fa-ba
ez-ih
ql-nq
kn-qo
yj-pv
cv-bj
yx-bg
ea-vo
vc-am
ee-zk
oq-tl
gb-yb
qh-hv
au-wj
wm-ot
xz-ob
af-ws
yd-qd
aa-yg
nq-uj
nh-qp
bn-jz
tx-fm
mh-vm
kv-ci
dm-zw
gn-fg
wc-tv
jg-py
wl-zj
zj-hk
rj-ze
el-tj
em-cv
qx-vg
gy-jw
fm-bq
bq-rg
wl-po
ex-cc
iw-pj
tn-hl
dj-tb
uc-ek
vh-mx
vo-oh
on-zy
by-bj
ur-un
bq-uy
lv-ce
bc-gd
op-oo
gz-tg
kk-hs
nv-dz
uj-ui
wm-vm
gz-rz
lm-ru
wb-zy
bc-vx
yt-jx
tw-am
ei-qw
sq-qx
nx-km
tq-nh
zk-tk
am-eb
tk-yj
mv-et
fs-kr
aa-oa
ik-ze
wd-yd
ta-oh
hu-bs
db-ot
om-ex
ik-rj
au-wf
hc-fx
km-hr
en-uu
zn-gk
yp-lx
gz-qf
cn-sa
ow-fh
zk-pv
er-ma
cj-nx
yc-zd
do-ty
hi-eb
ai-kl
jm-lh
qg-qk
he-wg
ld-xr
mj-ag
cs-ld
qj-ox
bs-sv
gb-zo
xi-wi
kt-mr
zo-bi
mv-nj
pr-ob
so-wy
xv-ja
om-cs
ho-zr
ov-dw
mk-sp
ar-hq
bd-xn
sd-dv
rq-sv
bl-zn
bw-kt
lg-tj
hm-dk
oe-sb
po-uv
pd-zy
rg-jm
bh-hc
co-ti
tb-cx
bl-df
wm-mh
pl-wm
qx-fy
wc-cl
wb-ts
dk-yn
tc-is
pc-iy
qd-tb
qj-ke
xe-ff
vv-ng
dq-fa
vo-nb
or-qh
wr-wq
fo-ji
jw-va
xm-ni
ex-cw
tv-zr
ex-ld
yb-bi
vb-ir
im-cq
bc-rl
an-ce
zg-xm
rb-hv
ih-hg
vc-wv
yy-kf
tu-tj
hc-jb
uo-yu
jl-dd
wn-wm
ju-ak
sz-pc
ox-ar
fh-th
jl-pf
nj-qf
gi-ha
gi-wv
lh-cy
om-cw
kx-yp
dl-iu
hd-pj
wi-zq
ab-tu
qg-lr
zn-yb
va-pq
yk-qj
xe-eu
mv-lu
xm-yl
if-xr
zh-fj
hh-lm
sg-qw
ac-kq
ek-tk
ft-by
xy-lg
iw-ak
is-ee
xs-mr
le-fr
iu-xa
wa-vg
bv-uv
cg-kq
mf-zy
oq-rj
vz-gg
tp-sm
kn-vp
op-dx
ub-vv
zt-qv
sm-xo
kd-kn
fj-ll
er-on
qk-cx
ua-kq
ok-yy
wk-af
ti-ar
ns-bc
ly-bb
hf-io
hf-ub
ze-eh
ra-vz
yk-hs
cv-uj
tp-ap
lo-ay
bz-ll
dq-bs
lt-oi
lj-en
wy-eg
pq-vt
on-sy
ge-xh
nk-au
bi-ff
ws-fp
sy-zy
eu-th
iu-cy
ua-bg
dg-nk
re-tv
to-ia
yx-cg
xb-zz
lj-sw
zb-rb
kg-xr
jl-im
cn-la
lm-fs
bx-he
pl-wn
rg-tx
pr-ez
ge-av
pw-yi
li-nv
nh-qm
lx-ce
aq-uo
yj-kp
pe-gb
eb-vm
dm-ql
ne-gh
tv-cl
mr-pg
sp-rx
ug-qj
cl-cx
no-et
aa-tq
ow-yo
hc-xk
fu-zc
dk-zi
wv-zb
th-ef
zr-dj
ij-vc
xv-iw
qf-hz
xd-qx
vb-qf
oo-zo
gy-vp
hf-fp
vx-kz
qo-ol
ad-vh
wj-wf
io-xm
cj-je
wj-kf
bv-cb
zn-fa
ze-oq
sp-oy
xb-qj
dz-up
qs-le
qj-ar
rf-is
dx-is
kg-sb
nj-no
rb-yo
qg-ok
yc-mb
mh-xp
ci-on
yb-sj
lj-bw
qm-aa
jl-bb
ov-vr
yr-hi
ik-uw
bx-ur
zy-ci
yu-le
ea-nk
qr-hc
jh-pi
hv-as
xp-wn
wb-tt
yo-ey
lg-tu
lv-nv
ns-wp
no-my
tm-jh
bc-qc
ml-rj
yi-rk
sy-cx
sw-ne
ws-gn
ez-rv
xp-eb
xa-dl
xb-mj
qm-lz
ur-xa
my-et
hr-rn
ft-cr
tm-zq
jt-br
ij-wv
yi-zq
au-sd
qo-jw
cr-hl
oa-nh
nq-xk
bg-us
cb-po
xb-ug
vr-oe
ya-we
ey-or
bt-vc
sn-fy
nt-jd
cw-cc
iv-az
zw-jb
eg-zi
ry-ac
jo-nf
lo-ha
im-ty
ja-mb
as-wx
iu-gl
qn-jm
oi-ld
gx-wx
rg-qn
wg-iu
ja-qz
un-nd
ei-lg
ms-pz
fz-ni
fp-fg
zf-qd
qh-rs
dh-ap
vg-ut
cg-ua
nh-aa
xn-sp
wv-am
df-oa
ei-ad
nc-tk
rf-mu
xz-ih
vx-rl
iy-jo
zs-sl
oh-wj
dz-yp
mb-qz
iw-zd
ag-ar
gl-oh
ip-pc
sv-qd
jd-do
tz-wl
xh-xd
ao-bv
pn-xl
fu-oi
tm-dp
kx-om
gg-hy
ge-wa
ww-dd
cy-wg
fh-ff
cw-cs
is-op
nl-ps
jp-oq
cb-hk
uc-cz
km-io
ns-xx
va-jg
ak-nm
aj-tb
im-bb
xf-vv
ty-pf
hg-un
ca-ud
ib-qc
uk-bx
pn-rn
aw-pn
ul-hr
zf-dq
xa-su
oy-fj
rk-vy
kn-gj
zb-or
jh-pw
wj-sd
zh-ln
qq-zj
yv-sa
ze-uw
fj-rx
ql-pt
ty-jl
io-fg
ee-nx
tb-hp
gz-lm
yb-ba
cy-ur
cs-iv
ky-ru
qh-ow
ih-pr
zb-ow
ub-bp
zi-qs
la-ra
tg-yf
qj-mj
vh-tj
bm-qo
nc-dx
yl-ws
mt-fh
dj-cl
pd-mq
wz-zc
tb-ho
lg-aa
tp-xo
nm-mb
vg-nj
mh-af
ns-kz
wh-af
qs-aq
qs-hm
gd-ib
jl-ww
tx-sz
cr-cv
tz-cb
am-ay
fg-td
nl-gy
wl-zp
xb-yk
st-kv
oz-ay
ma-ts
pk-yf
bc-wp
sv-sl
yb-fa
lj-pg
ra-hy
sf-bj
te-rj
if-dw
hk-uv
jb-mu
uv-zj
af-wm
ke-uz
je-pn
cl-gq
sw-pg
ry-hw
pl-yr
lo-bt
zo-sj
lj-uu
tt-kv
ns-fd
kd-it
xr-ud
ij-ha
et-to
jt-xk
qn-uy
ps-qo
xk-zw
hs-ug
wj-ta
di-my
xo-dh
po-qq
hd-zd
ln-gr
oj-kd
kz-os
wr-hu
xi-nn
gi-ox
mj-ug
fm-rg
jf-ex
jl-jd
qp-oa
qr-jb
dq-zs
jp-rj
el-vh
le-eg
yn-yu
lv-dz
jw-oj
ze-te
rm-uo
we-ih
np-an
xi-tm
ib-vx
iy-jm
re-wc
hl-uj
ws-zg
fj-mk
kg-ud
zm-zp
mu-zw
sa-ek
zk-nc
zz-jg
bd-fi
qc-xx
ma-zy
an-yp
vr-ji
nb-oh
ez-jv
bd-rx
li-dz
jf-tq
tq-lz
uz-hs
aw-xl
qo-it
qs-eg
op-pv
in-os
vp-ps
dj-gq
kn-jw
sg-el
on-ts
qo-gj
oa-bf
br-hc
pt-nq
kr-pk
xf-pz
fp-yl
up-lx
hh-xm
ze-vy
jc-lb
gd-rl
pe-oo
ad-xy
fa-db
qd-rq
co-th
jb-nq
sl-wd
ok-pz
uw-kh
eq-zz
cl-qk
wp-ib
yk-ag
sg-ei
np-up
gz-fs
qv-tq
wv-ha
rx-zh
xz-ro
dv-wf
bm-oj
ol-it
wg-lh
ks-ov
bo-sn
bz-mk
hp-cx
nv-an
qc-kr
nd-iu
dh-ix
nf-fm
nj-et
gl-sd
ql-jb
qs-yn
fu-cs
yf-ky
mq-kq
zb-ey
kd-nl
br-dm
kp-nc
le-rm
xf-lr
cz-vz
wv-lo
xi-rt
av-mx
sd-ea
wp-xx
lm-id
mv-di
gq-aj
kk-hq
zk-rf
zp-qq
jm-bq
tu-sg
kk-ar
uv-zm
uy-jm
ip-jm
kn-te
dk-aq
ho-gq
bw-sw
mq-cg
fo-ca
gy-gj
lt-ro
yc-iw
zf-wr
eq-jg
ou-am
av-xh
of-et
sj-gk
uc-ko
fr-ai
ag-ke
ae-el
ww-nt
gn-fp
ij-lo
tb-wc
hi-pl
iu-ur
mh-pl
hg-ya
dp-of
wd-zs
hw-bg
hc-zw
ry-mq
an-up
vv-ms
if-oe
ly-im
zt-bf
xx-rl
wf-rz
xy-qw
jx-km
ky-hh
uw-vy
ts-sy
jz-xz
cg-bg
rm-eg
qw-tu
lm-rz
of-nn
fg-ni
th-ti
rx-gr
tu-mx
ba-kk
tl-ml
gd-kz
ix-pq
wa-xh
fm-sb
iv-fc
fs-cg
db-zo
ha-vc
yi-xi
vj-ay
dx-tc
qz-ju
sf-jj
xe-yw
mf-er
rl-oz
xn-zt
ti-ef
hq-ug
ok-lr
of-tm
qz-hd
ot-pl
io-ws
fo-ks
rq-wr
lt-iv
lv-up
ex-oi
dj-wc
vo-sd
hn-iu
ln-rx
jm-sz
ia-lu
jb-fx
uc-ra
qc-fd
ro-rv
ox-uz
ke-hs
rs-gx
yy-gu
sp-gr
et-vb
ol-ps
ta-dv
tu-vh
ce-np
ya-ez
lr-qj
aq-yu
hn-nd
di-qf
ov-xr
wa-sq
lh-dl
eq-sm
kd-vp
tp-pq
lp-cg
oz-ou
gg-ko
kv-er
dw-vr
nd-he
cz-ko
uw-eh
kl-bw
kt-en
ly-do
xa-un
hz-et
tm-rt
by-tn
hu-wd
hh-rz
rq-wq
wn-hi
jp-kh
pd-us
zz-xo
wm-hi
ba-sj
kg-ty
jv-pr
jh-zq
su-ur
ne-kt
dx-rf
tk-rf
dv-ea
lu-et
ke-kk
eb-ot
rb-rs
st-on
pt-xk
nk-oh
tj-xy
ds-tn
ei-mx
bz-oy
ba-oo
xl-yt
lb-li
yk-ug
ko-cn
vv-sm
om-zc
qj-hq
eh-te
qe-tl
ge-fn
fs-hh
ee-pv
ru-id
un-hn
mv-my
uk-as
rj-vy
hf-zg
hn-ur
pj-yc
aq-le
wb-ma
sg-ad
it-ps
wt-cb
jl-ls
ci-gc
ba-pe
co-ft
wd-sv
hs-sw
qd-zs
vx-os
zy-gc
ea-nb
wl-bv
lp-ac
ni-td
nj-my
wr-zs
tz-zj
ui-sf
kd-jw
hy-cn
ol-nl
zf-sv
xp-hi
dm-nq
uc-bh
yd-ay
qw-ae
hi-wh
hx-ff
cr-bj
bj-ds
us-lp
wt-zp
uk-rs
ap-vt
xz-jv
gb-db
oq-qe
yo-as
ry-kq
ol-bm
yn-zi
km-rn
ia-et
ik-ml
nh-jf
qq-cb
aw-zn
yc-ak
ap-eq
ob-jz
gi-vj
lt-wz
rz-pk
lo-vc
is-yj
lv-jc
vr-ks
ff-yw
rv-bn
gk-ba
tn-wd
nm-tv
gh-lj
di-nj
we-pr
yw-qi
aq-zi
ra-cz
bn-ya
aw-nx
vt-jg
dl-su
kh-te
ow-as
mx-el
mv-to
ja-nm
tk-ee
ta-wf
pl-kv
hs-ox
hz-nj
bg-ry
ln-fi
xb-ar
tp-va
jf-lz
le-hm
bp-lr
jh-rt
yy-ub
zy-kv
gz-kr
yv-ra
pt-jb
kp-pv
ai-lj
rb-or
ip-kx
bw-uu
gx-hv
dd-ty
zm-ao
tt-ok
hv-yo
cz-sa
us-ua
pj-qz
id-yf
aj-cx
df-lz
aq-so
mh-ot
uo-hm
ru-pk
mt-co
qw-cq
ek-vz
sm-jg
cv-tn
cq-do
ea-ts
kq-us
lw-bg
wf-sd
ho-tv
qr-ql
nd-ur
lx-lv
dd-pf
qs-uo
if-kg
uu-kt
oj-ge
rs-wx
lw-us
ja-ju
ti-hx
gb-bl
zs-zf
yv-vz
ry-lw
dp-xi
zi-rm
nj-ia
ge-qx
hp-wc
yf-hh
ng-xf
ty-bb
we-bn
gj-it
zh-ll
dg-wf
zk-yj
lt-cw
hy-yv
hp-re
ks-sb
gh-sw
wy-yu
wr-sv
nt-pf
zg-fc
zk-dx
gl-nk
ke-wr
zv-hh
dk-le
jh-rk
rq-bs
kf-qg
ql-mu
ul-km
wh-vm
dk-eg
oo-gb
bb-cq
mq-hw
nq-br
xe-vi
qn-sz
fh-ti
xx-gd
bt-ij
zt-lx
wn-eb
ou-ay
jd-im
xi-rk
xi-jh
ov-if
gu-ub
tu-ae
py-tp
jm-fm
sp-zh
iv-cw
dh-py
jj-bj
bw-pg
eq-vt
ij-am
ky-id
xv-yk
hu-sl
ob-rv
gq-wc
bt-tw
ro-bn
pv-dx
dz-bz
qi-ti
fz-fg
dx-yj
nx-jx
qr-jt
ox-ag
yr-mh
bf-qp
yj-xs
vj-jg
pd-ua
jt-mu
rs-yo
nt-cq
qk-aj
uo-so
ur-wg
ik-eh
cb-zp
rm-qs
ra-gg
pt-ru
la-bh
mb-pj
pe-db
fm-iy
ir-hz
lj-ne
qp-jf
hs-xb
kt-gh
to-ir
ha-tw
zo-bl
ef-mt
ce-kx
di-vb
rg-uy
kl-kt
ln-ll
dd-ly
ls-bb
li-kx
mt-ti
nb-au
kz-wp
ai-cj
ls-ww
dg-nb
zz-dh
mq-lw
wb-la
pd-cg
aw-ul
kh-eh
ez-jz
pi-dp
fu-di
rz-fs
nf-iy
zt-tq
ui-cv
xa-zv
bx-cy
yv-bh
vv-kf
bh-vz
nh-bf
tc-uz
ag-ey
dx-tg
cv-ds
yn-uo
wx-uk
cg-hw
ik-vy
vi-th
rq-sl
sl-bs
ci-br
fh-yw
ob-ih
zb-yo
fz-yv
ml-qe
ut-ge
jv-hg
uz-mj
dl-bx
sy-ci
bp-vv
wn-vm
ij-gi
pe-sj
kk-uz
kk-xb
ll-mk
io-fc
hi-af
fr-gh
it-bm
yy-ms
kr-hh
yf-rz
bj-ui
cz-la
oh-dv
cj-pn
ui-ds
po-zp
cz-yv
kt-lj
uq-nm
tk-xs
id-pk
wk-yr
ja-zd
bt-am
qh-ey
ac-cg
kf-ok
ud-ks
rb-tc
aa-df
ui-cr
on-cv
oe-xr
eb-pl
dh-pq
kz-he
ac-yx
hk-po
te-oq
xx-vz
ce-li
zn-db
bp-kf
ui-tn
ty-ls
jx-pn
xl-hr
lv-ca
tv-gq
tv-aj
wq-yd
mr-en
ip-iy
yr-wh
xr-ca
jz-ih
yw-hx
lb-nv
sv-wq
ls-pf
kd-gy
wm-wh
wx-rb
ao-po
sl-yd
qp-aa
ns-qc
az-zc
uk-ey
in-ns
ds-uj
ww-ty
pw-zq
eu-ff
wd-zf
uq-iw
er-sy
qr-fx
po-zj
io-gn
ln-lu
yn-hy
ag-qj
oa-qm
bc-gb
zd-uq
jc-rj
yv-la
rm-dk
ln-fj
jf-aa
cj-aw
if-ji
dk-qs
mu-br
ea-gl
gi-am
we-sv
dw-ks
in-jt
jd-qv
wm-eb
oq-kh
hv-ow
uw-oq
ji-ca
bv-zm
iy-tx
in-bc
yp-nv
nk-wf
fh-ef
km-pn
ab-ei
dh-wy
mf-sy
gb-zn
qn-nf
dl-wg
zr-wc
wc-aj
cq-ww
ws-td
oj-vp
hw-lw
jp-tl
dm-pt
mf-tt
xa-ha
ll-rx
yu-qs
vz-hn
wk-mh
df-nh
un-he
pq-iw
qc-gd
ly-ls
bf-nf
vy-te
nn-yi
ob-ro
gx-or
nb-sd
vy-kh
gc-tt
zq-ov
yp-ce
yu-so
qm-jf
pi-yi
mj-qs
sn-fn
dl-he
os-ib
gy-qo
wr-qd
sn-av
hi-wk
qm-bf
bj-hl
sq-av
oz-gi
cs-lt
kh-ml
ae-ad
hf-ws
kf-ub
co-qi
rl-fd
we-ro
nt-ty
ox-uj
as-rb
fs-pk
rt-wi
jd-ls
bf-qv
ds-hl
ng-pz
ra-sa
vt-va
fd-oh
bg-au
ba-zn
gg-uc
wi-tm
hk-bv
vj-so
uk-ow
xx-kz
yv-ek
vt-zz
qs-wy
bb-ww
eu-hx
zg-ni
gg-sa
pe-bl
xp-af
rg-iy
eh-vy
pr-yk
bq-iy
ip-qn
rv-zk
jj-by
jc-nv
vo-au
sa-uc
yg-df
no-mv
ij-xn
ml-eh
oe-kg
qq-tz
gi-lo
ry-pd
th-hx
vt-ix
gd-in
wt-qq
ud-ji
hn-dl
mt-xe
yt-ul
vh-lg
eq-ix
eb-mh
rx-oy
fp-xm
os-bc
ky-rz
bh-gg
dp-zq
sq-bo
au-ta
jh-of
ao-uv
wn-af
vg-sn
ei-tj
lt-om
id-gz
bs-qd
oe-jp
ac-bg
ru-hh
nk-wj
bz-xn
qr-xk
mh-wh
dl-cy
qf-my
ia-my
oa-lz
yp-up
cc-oi
qm-df
mk-rx
wl-ao
el-ei
ok-ub
mb-iw
on-tt
ap-sm
hm-zw
ov-sb
py-vt
nc-tc
rv-jz
as-vh
mk-gr
cv-by
th-xe
hf-ho
ag-hq
wg-un
ex-lt
jv-ya
yr-af
sf-dv
ly-jl
yj-tc
it-gy
zz-sm
ze-ml
xl-km
cq-ty
us-fg
qv-aa
mr-uu
bx-hn
xz-ya
np-jc
bn-ob
uu-pg
hy-sa
eu-gu
tq-nm
sb-ji
ja-iw
wx-xv
ks-if
ul-jx
xe-ti
mj-ke
py-zz
gc-wb
rm-aq
bj-em
vb-mv
xs-nc
wj-nb
qg-bp
gh-kl
xn-ln
ox-xb
oj-kn
ff-th
ab-zh
hq-dd
uj-em
jx-ur
cn-vz
ke-xb
sg-tj
rl-qc
va-xo
zp-ao
bb-jd
ty-jd
bl-sj
cn-vo
jx-hr
dj-hp
jf-zt
zw-br
pk-tg
jj-cv
nx-pn
hw-eq
ln-sp
ek-cn
xv-ju
im-mk
qg-yy
jb-br
av-xd
nf-jm
mq-bg
eg-yu
kh-rj
hi-ta
eh-qe
po-wt
hp-aj
cn-ra
op-zk
fa-bl
pe-gk
ro-jz
jx-aw
ab-sg
jp-uw
gu-ok
zh-gr
kf-ms
ib-ns
gc-vi
bo-ut
zm-sn
ji-qm
cc-ld
ex-zc
cj-xl
hn-he
le-zi
ru-fs
eu-co
yw-eu
df-qv
hm-wy
sj-gb
dh-ds
kq-hw
jp-ml
bh-ko
xy-je
ut-qx
ij-ou
ov-oe
st-mf
pv-is
bp-ok
ww-pf
rs-zf
in-xx
br-qr
ia-hz
zo-ba
nh-lz
zn-pe
rk-zq
xl-je
dq-hu
bq-tx
oe-ks
tj-qw
uj-bj
uy-jo
dh-va
lt-az
dg-vo
ud-fo
je-rn
wl-uv
ih-bn
zr-qk
jv-bn
wt-ao
wp-qc
fs-id
fr-sw
fd-bc
vj-bt
dm-jt
ci-mf
pd-lp
yn-aq
av-wa
dq-yd
op-nc
ty-ly
wa-bo
bw-fr
ft-hl
fr-lh
wa-fy
uq-ak
bi-oo
xf-ui
wx-yo
ae-tj
uo-dk
bv-wt
qe-vy
lz-aa
zk-kp
cn-bh
ad-lg
ms-lr
fi-oy
hr-hc
hc-dm
fx-dm
dp-wi
mk-bd
qv-lz
ld-iv
do-pf
oj-gj
pz-gu
em-by
tb-cl
jf-qv
dw-kg
ak-zd
wq-mt
dm-zr
sl-st
ov-kg
uw-qe
dd-kd
ix-ap
sy-kv
hm-eg
cz-cn
dj-qk
do-ls
lb-lx
wv-tw
nt-ly
bv-qq
gn-yl
ge-fy
em-cr
dw-oe
sw-kt
ms-bp
km-je
hf-td
om-cc
sn-ut
lu-di
xz-we
st-wb
va-eq
py-pq
au-gl
pd-ac
vc-tw
le-so
uv-wt
lo-tw
fm-qn
ge-xd
lb-an
aw-yt
bs-wd
ar-yk
hh-tg
kn-nl
su-hn
xs-dx
iw-ju
cc-lt
gk-bt
qp-yg
qw-lg
xv-pj
cz-ek
fu-ex
qw-el
tq-yg
zb-wx
dp-jh
xm-gn
io-ni
jh-nn
yl-zg
oh-au
em-ds
jm-pc
qm-zt
ms-xd
ea-wf
vj-am
ps-kn
gz-hh
nf-ip
zr-re
pj-ak
wk-ot
oq-ml
ot-wn
nn-dp
yl-hf
tx-nf
aj-re
pk-gz
qp-qm
wo-zi
xr-sb
zm-zj
qf-mv
ya-pr
lx-nv
dg-dv
va-sm
oz-vj
tp-eq
an-lv
gc-st
ci-tt
wq-wd
iu-lh
ni-yl
pi-zq
vb-ia
hd-nm
oy-ll
ld-fu
yn-wy
sa-lo
om-az
qp-qv
rn-yt
oh-sd
kx-lv
aj-zr
to-lu
wp-os
qw-vh
zv-lm
bw-ne
rt-yi
by-hl
fn-wa
dj-aj
ay-gi
yb-bl
iy-sz
mr-kl
gx-qh
ae-ze
ak-hd
of-rk
ad-mx
ce-jc
fh-qi
ol-vp
wp-in
ni-gn
we-jz
ln-mk
ac-us
ra-ll
ru-kr
vm-vx
tm-rk
hm-so
jw-it
qe-ze
cy-xa
fn-vg
ia-ir
vm-yr
yr-wn
je-ul
pi-pw
kr-tg
bs-zf
sq-fn
vr-ca
fp-ni
ma-mf
jp-te
pl-xp
ob-we
bt-ha
nf-uy
hd-iw
ov-ud
fa-sj
su-un
pf-ut
eg-so
br-xk
vb-no
qc-kz
br-fx
vb-nj
yx-ry
bz-zh
ja-yc
oz-tw
ap-zz
kf-lr
cg-lw
zs-wq
gd-os
gc-ma
wh-pl
bm-nl
ar-ug
zi-so
yl-td
ow-or
yx-hw
ji-xr
rb-uk
cx-zr
zj-gj
ju-mb
ib-rl
tm-np
kv-ts
nf-rg
xv-ak
ly-pf
ql-fx
kp-rf
qx-av
tp-vt
jx-rn
ky-tg
ad-ab
hk-fj
yc-uq
wx-hv
su-is
yi-dp
sj-bi
wo-km
xm-hf
hw-us
df-ir
fz-gn
lz-zt
gy-vh
jj-hl
to-no
mu-fx
qo-no
fa-zo
hg-jz
lx-np
em-hl
bx-nd
qp-df
ob-ya
mr-fr
ng-ok
is-zk
tn-cr
tl-ik
eu-fh
lz-yg
pc-rg
xs-op
xv-yc
bp-xf
mf-kv
zi-yu
lh-nd
oh-wf
re-qe
gn-sd
tz-bv
uy-pc
ji-ov
fr-pg
ou-gi
uz-ug
ak-qz
bj-tn
wg-oo
yc-nm
mb-ak
nl-it
wi-nn
id-zv
hr-yt
fp-zg
wa-li
hv-or
lh-ur
ff-mt
hf-fz
ta-vo
yb-zo
uu-gh
ja-ak
qn-jo
np-dz
nd-wg
cz-hy
nk-dv
an-kx
db-sj
nx-hr
os-ns
qo-nl
so-yn
eu-vi
dl-fx
vr-sb
nn-zq
ix-sm
jp-ik
yb-db
bj-lm
ey-wx
le-wy
jo-rg
kv-on
nj-to
fi-bz
ww-lb
ub-qg
us-mq
zi-wy
er-zy
to-di
pj-nm
uv-vr
gd-kq
bq-jo
uy-tx
oa-zt
yp-lb
xm-fz
zd-pj
pt-qr
db-aq
yf-zv
bx-su
or-as
ql-hc
as-rs
no-lu
uw-te
oy-fy
vj-tw
ft-jj
eq-pq
df-bf
dz-jc
gk-zo
oj-ps
bi-pe
ok-xf
yl-fz
sm-ih
hr-wo
bf-jf
eh-tl
cy-un
xk-jb
bw-en
nf-sz
xd-sq
sa-la
ze-uy
ow-wx
ab-mx
hu-zf
ke-yk
nt-bb
wd-dq
gn-td
zb-gx
yg-oa
xf-qg
oy-gr
ma-tt
cb-ao
mq-lp
ea-dg
kp-lo
af-vm
ci-ma
cl-ho
xi-pw
ir-my
mb-uq
rz-zv
gi-tw
ob-hg
aq-hm
ms-ub
ju-zd
yt-hy
wf-gl
db-gk
pq-zz
lg-ab
rm-so
wb-kv
jg-pq
bp-bw
hd-ju
hr-cj
cy-he
gj-ol
qg-pz
vi-fh
qd-wd
wg-bx
tp-jg
if-ud
mr-lj
st-ts
xz-rv
lr-yy
py-xo
td-zg
tb-zr
yb-pe
vi-co
tb-tv
nc-yj
tn-em
rj-tl
wm-xp
bx-xa
ix-jg
fr-kt
tk-kn
fj-sp
wy-rm
ln-oy
zb-qh
hc-nq
hl-wc
yk-hq
hs-mj
ry-ua
jb-ac
ha-ay
fd-xx
yp-li
ar-ke
tc-kp
pz-bp
my-lu
zt-qp
an-lx
ro-jv
fo-if
fa-bi
gr-fj
qm-tq
tq-oa
qf-lu
cn-gg
az-ld
lm-pk
mb-hd
iy-ad
vo-gl
xo-qh
qg-ng
ub-lr
of-rt
if-ca
hs-qj
yy-xf
xb-ag
ua-mq
bp-yy
mt-qi
bn-ez
yn-rm
bz-gr
uz-qj
ls-cq
mr-ai
sn-ge
rb-qh
gr-fi
kh-tl
hp-tv
gr-eh
ot-xp
jp-qe
so-dk
iu-un
in-ib
cw-wz
bn-pr
cl-re
sn-xd
xk-ql
hv-en
tc-zk
hk-zm
tz-zp
hp-ho
ps-gj
ik-qe
of-pi
ey-rb
yk-uz
ho-dj
fz-fc
sd-ta
la-ek
ut-fn
av-ut
sj-zn
dh-vt
re-qk
in-fp
uo-le
mx-ae
ay-tw
ds-cr
xm-td
ry-ak
lt-ld
gx-uk
wv-bt
hq-uz
cl-ug
er-tt
wv-ou
lr-vv
tc-xs
pz-lr
ek-ra
kp-op
fr-lj
ya-yb
wp-rl
ov-ca
bd-bz
yp-el
gu-bs
ql-jt
gu-ng
no-di
xf-gu
nl-yg
sa-vz
bo-ml
lp-ua
dd-im
yl-fc
rv-hg
ee-kp
ow-ey
ca-sb
er-wb
uk-qh
to-my
jo-tx
lb-ce
ai-uu
ap-xo
my-up
zd-yl
xe-qi
xf-ms
mj-hq
wy-aq
sq-sn
qj-kk
fg-fc
db-bl
wk-xp
fn-av
zv-kr
vi-kl
yd-wr
dg-sd
zp-hk
jj-ds
gz-zv
bm-jw
lg-el
ne-mr
xr-dw
ta-nb
vc-gi
kl-sw
dk-wm
jg-xo
wq-dq
wb-mf
ze-jp
to-hz
xh-ut
ab-el
bi-gb
oi-sq
nc-is
aa-zt
zv-ky
tl-vy
xl-jx
kd-ol
gx-ow
zg-fg
do-bb
hl-cv
hg-bn
fd-os
gg-cz
kv-ma
sf-cr
hs-hq
pt-mu
oj-ol
mq-ac
vv-ok
mk-fi
wl-so
et-qf
wh-ot
cr-by
dl-ur
iv-cc
tl-jl
xi-pc
hz-no
pi-rt
hi-ot
dk-yu
lo-oz
ef-ff
vg-xh
oz-am
wn-wk
wk-eb
cb-zj
pf-jd
ws-ni
pv-tc
ud-dw
uu-ne
pf-cq
qi-vi
bd-zh
kx-nv
qw-mx
rt-zq
qm-qv
ez-hg
xo-vt
sb-ud
mb-zd
oj-qo
st-tt
it-oj
tx-ip
uj-sf
zo-zn
qz-zd
zm-po
mx-tj
qh-yo
ap-va
jc-yp
uy-ls
bl-oo
wf-xk
lz-bf
wa-sn
pd-lw
to-vb
bd-ps
ik-te
pe-fa
fy-xd
ky-lm
un-bx
sl-dq
pk-ky
ne-ai
dg-wj
wp-gd
ai-en
yx-ut
vv-pz
pi-xi
hl-ui
xx-os
lg-ae
gl-ta
rf-yj
uc-hy
im-ls
wt-wl
pv-xs
fr-ne
ur-he
ot-yr
ek-ko
oe-fo
wq-qd
lh-bx
wi-rk
cy-hn
nt-do
ai-sw
xd-vg
jd-cq
ps-bm
to-qf
dv-wj
bz-fj
tn-sf
fu-lt
en-ne
yw-co
yk-mj
vg-sq
yy-pz
yv-ko
pn-vc
ry-us
er-st
bf-tq
vg-bo
xv-qz
ea-wj
dw-fn
gh-pj
wy-uo
ao-hk
ky-kr
jd-ly
mb-xv
jo-pc
np-lb
tj-cs
hg-we
nd-dl
ns-vx
wm-wk
bd-gr
ro-pr
hp-ez
rz-ru
qf-ei
bm-gj
iv-zc
yi-jh
pz-kf
ok-ms
vj-ha
ky-cz
af-pl
tt-ts
av-fy
ia-di
dq-sv
yo-uk
ll-bd
vb-hz
yl-io
dv-au
oh-dg
bq-pc
bq-vp
cg-ry
ek-bh
dh-eq
cs-oi
bh-ra
ox-ug
qd-hu
sp-fi
li-lx
id-tg
sp-bz
oz-bt
hv-rs
oi-om
yd-zf
jd-ww
uk-zb
bw-gh
iv-wz
wl-qq
xa-he
kf-ng
pt-hc
pw-oa
te-ml
oa-jf
tz-ao
qf-ia
ca-dw
rg-az
yr-id
io-td
qg-ms
by-sf
uq-hd
jw-gj
mt-eu
dm-xk
td-fz
ll-xn
vb-my
ea-oh
vg-av
qz-yc
mv-ir
ud-vr
tp-ix
jf-yg
mt-th
hw-ua
nm-qz
wo-yt
th-qi
hc-jt
qx-wa
sd-nk
ol-nv
aq-eg
im-do
je-hr
gz-ru
su-nd
xf-ub
gk-bi
bt-ay
kn-gy
qh-as
je-nx
yv-cn
yy-qz
rn-ul
xh-bo
fo-dw
hg-pr
xe-fh
el-tu
mh-pv
ae-vh
hm-yu
xz-ez
bb-dd
qx-sn
id-rz
mv-hz
ix-xo
kh-to
by-uj
bq-qn
qo-vp
us-cg
wo-pn
xv-zd
ll-sp
ut-fy
vi-ef
hd-zs
ay-vc
hn-xa
ek-hy
hq-ke
hx-co
aj-ho
xf-kf
cc-zc
jj-em
xm-fc
sz-rg
tl-te
mu-qr
op-tc
su-iu
qm-yg
yf-ru
iw-qz
bd-fj
xb-hq
wn-wh
bp-gu
lb-kx
zz-tp
je-yt
iy-ao
id-hh
zr-hp
em-ft
ui-by
pd-kq
em-ui
mj-ox
qk-gq
sq-xh
hx-fh
dg-uq
tk-kp
kl-ne
sm-pq
sz-uy
yn-le
ge-vg
fc-qr
zp-uv
or-yo
fo-rx
ro-ez
bs-zs
sq-ge
dq-rq
fn-xd
xa-nd
zo-gx
hf-gn
uj-jj
sg-mx
tb-gq
wr-bs
fr-en
we-ez
nb-ly
rl-er
tm-nn
am-ha
bv-jj
hy-la
pt-br
wt-zj
if-vr
yt-km
cv-ft
lv-li
ek-gg
lh-hn
yg-zt
yx-pd
ef-hx
zh-mk
wi-yi
fo-kg
ma-ja
zc-bm
ma-sy
ft-bj
zh-fi
vh-ei
fi-rx
xn-fj
uq-xv
nx-nk
ni-fc
ua-sv
up-lb
fs-tg
sf-em
cx-re
bd-oy
tz-wt
xr-vr
dd-jd
zn-bi
cl-aj
do-jl
nq-fx
oz-ij
sz-bq
bz-rx
li-an
li-jc
ua-lw
rs-or
xe-ho
ou-kt
zp-zj
pd-bg
tk-dx
fa-gk
bq-ip
dd-ls
ir-no
po-bv
qr-zw
wb-on
hr-aw
yx-us
fs-wc
yx-ft
qk-tb
tx-pc
fd-vx
xn-mk
zv-ci
do-ww
rn-nx
fx-xk
cr-uj
tz-zm
vi-ti
ug-ke
fd-gd
lw-kq
tt-zy
sf-hl
cj-rn
vm-hi
as-zb
fd-ib
bw-ai
nc-rf
qg-gu
mb-jz
xh-qx
yf-rn
oz-ha
pg-py
st-ci
jv-we
fo-ov
lp-bg
wr-wd
jt-jb
uu-fr
ft-tn
ix-py
ib-xx
mv-ia
fs-yf
oz-vc
sw-uu
xs-is
ij-ay
nm-iw
ks-xr
eh-rj
rf-ee
ll-fi
fx-jt
rf-xs
wv-vj
zm-dp
bl-ba
qq-uv
tg-zv
yw-th
jo-jm
he-su
ex-iv
az-cs
dp-pw
bo-xd
sg-vh
wb-sy
ng-ub
ex-az
nv-up
ix-ef
dg-gl
ks-kg
ho-cx
ro-ih
db-bi
gl-nb
nh-zt
ad-tj
bo-qx
ju-by
pl-wk
su-cy
ef-eu
zp-bv
bi-ba
np-li
ju-nm
om-fu
qe-kh
fx-pt
qi-hx
ho-qk
wo-aw
xy-sg
nn-rk
rk-rt
ge-bo
rq-wd
pg-kt
wi-pw
ee-yj
lr-bn
gh-ai
jx-je
la-gg
qe-te
ib-bc
ko-la
ce-nv
dj-re
ab-qw
rk-pi
sq-fy
pe-ne
ip-uy
vx-qc
wc-cx
ee-nc
wl-hk
dm-mu
hi-mh
uq-qz
ff-qi
wi-wz
zg-fz
ot-af
ji-kg
lw-lp
ox-ke
zg-ts
pw-rk
pf-im
gj-kd
wv-oz
hz-di
wn-as
tm-pi
jw-nl
ff-co
yf-gz
yc-ju
yi-of
np-kx
zn-oo
tg-rz
as-gx
os-rl
nl-oj
ce-ql
kg-vr
bf-yg
wp-eg
kl-hu
ko-sa
po-tz
fc-ws
zc-oi
kl-pg
vz-uc
cs-ex
xi-zq
lv-np
pg-en
fj-fi
td-fp
im-nt
el-ad
lm-kr
tv-qk
jv-jz`);

/**
 * the code of part 2 of the puzzle
 */
const part2 = (input: string) => {
  const graph = input
    .trim()
    .split("\n")
    .reduce<{ [key: string]: string[] }>((obj, line) => {
      const [left, right] = line.split("-");
      if (obj[left] === undefined) obj[left] = [];
      if (obj[right] === undefined) obj[right] = [];

      obj[left].push(right);
      obj[right].push(left);
      return obj;
    }, {});

  // the maximum connection size should be the one where all connections are connected to each other
  const MAX_LENGTH = Math.max(
    ...Object.values(graph).map((array) => array.length)
  );

  // find the biggest set intersection between all connections
  let biggest: string[] = [];
  Object.keys(graph).forEach((node) => {
    const possible = combination(graph[node].sort(), MAX_LENGTH - 1);
    for (let i = 0; i < possible.length; i++) {
      let common = new Set([node, ...graph[node]].sort());
      for (let j = 0; j < possible[i].length; j++)
        common = common.intersection(
          new Set([possible[i][j], ...graph[possible[i][j]]].sort())
        );
      if (common.size === MAX_LENGTH) biggest = Array.from(common);
    }
  });

  return biggest;
};

export { part1, part2 };