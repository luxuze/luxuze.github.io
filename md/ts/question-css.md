# CSS

## css sprite是什么,有什么优缺点

1. 概念：将多个小图片拼接到一个图片中。通过background-position和元素尺寸调节需要显示的背景图案。
1. 优点：
    1. 减少HTTP请求数，极大地提高页面加载速度
    1. 增加图片信息重复度，提高压缩比，减少图片大小
    1. 更换风格方便，只需在一张或几张图片上修改颜色或样式即可实现
1. 缺点：
    1. 图片合并麻烦
    1. 维护麻烦，修改一个图片可能需要从新布局整个图片，样式

## display: none;与visibility: hidden;的区别

1. 联系：它们都能让元素不可见
1. 区别：
    1. display:none;会让元素完全从渲染树中消失，渲染的时候不占据任何空间；visibility: hidden;不会让元素从渲染树消失，渲染师元素继续占据空间，只是内容不可见
    1. display: none;是非继承属性，子孙节点消失由于元素从渲染树消失造成，通过修改子孙节点属性无法显示；visibility: hidden;是继承属性，子孙节点消失由于继承了hidden，通过设置visibility: visible;可以让子孙节点显式
    1. 修改常规流中元素的display通常会造成文档重排。修改visibility属性只会造成本元素的重绘。
    1. 读屏器不会读取display: none;元素内容；会读取visibility: hidden;元素内容

## link与@import的区别

1. link是HTML方式， @import是CSS方式
1. link最大限度支持并行下载，@import过多嵌套导致串行下载，出现FOUC
1. link可以通过rel="alternate stylesheet"指定候选样式
1. 浏览器对link支持早于@import，可以使用@import对老浏览器隐藏样式
1. @import必须在样式规则之前，可以在css文件中引用其他文件
1. 总体来说：link优于@import

## display,float,position的关系

1. 如果display为none，那么position和float都不起作用，这种情况下元素不产生框
1. 否则，如果position值为absolute或者fixed，框就是绝对定位的，float的计算值为none，display根据下面的表格进行调整。
1. 否则，如果float不是none，框是浮动的，display根据下表进行调整
1. 否则，如果元素是根元素，display根据下表进行调整
1. 其他情况下display的值为指定值
1. 总结起来：绝对定位、浮动、根元素都需要调整display

## position的值， relative和absolute定位原点是

1. absolute：生成绝对定位的元素，相对于 static 定位以外的第一个父元素进行定位
1. fixed：生成绝对定位的元素，相对于浏览器窗口进行定位
1. relative：生成相对定位的元素，相对于其正常位置进行定位
1. static 默认值。没有定位，元素出现在正常的流中
1. inherit 规定从父元素继承 position 属性的值

## display:inline-block 什么时候不会显示间隙？(携程)

1. 移除空格
1. 使用margin负值
1. 使用font-size:0
1. letter-spacing
1. word-spacing

## 清除浮动的几种方式，各自的优缺点

1. 父级div定义height
1. 结尾处加空div标签clear:both
1. 父级div定义伪类:after和zoom
1. 父级div定义overflow:hidden
1. 父级div也浮动，需要定义宽度
1. 结尾处加br标签clear:both
1. 比较好的是第3种方式，好多网站都这么用

## display有哪些值？说明他们的作用

1. block 象块类型元素一样显示。
1. none 缺省值。象行内元素类型一样显示。
1. inline-block 象行内元素一样显示，但其内容象块类型元素一样显示。
1. list-item 象块类型元素一样显示，并添加样式列表标记。
1. table 此元素会作为块级表格来显示
1. inherit 规定应该从父元素继承 display 属性的值

## 介绍一下标准的CSS的盒子模型？低版本IE的盒子模型有什么不同的

1. 有两种， IE盒子模型、W3C盒子模型；
1. 盒模型： 内容(content)、填充(padding)、边界(margin)、 边框(border)；
1. 区 别： IE的content部分把 border 和 padding计算了进去;

## PNG,GIF,JPG的区别及如何选

1. GIF
    1. 8位像素，256色
    1. 无损压缩
    1. 支持简单动画
    1. 支持boolean透明
    1. 适合简单动画
1. JPEG
    1. 颜色限于256
    1. 有损压缩
    1. 可控制压缩质量
    1. 不支持透明
    1. 适合照片
1. PNG
    1. 有PNG8和truecolor PNG
    1. PNG8类似GIF颜色上限为256，文件小，支持alpha透明度，无动画
    1. 适合图标、背景、按钮

## 行内元素float:left后是否变为块级元素

1. 浮动后，行内元素不会成为块状元素，但是可以设置宽高。
1. 行内元素要想变成块状元素，占一行，直接设置display:block;
1. 如果元素设置了浮动后再设置display:block;那就不会占一行。

## 如果需要手动写动画，你认为最小时间间隔是多久，为什么？（阿里）

1. 多数显示器默认频率是60Hz，即1秒刷新60次，所以理论上最小间隔为1/60＊1000ms ＝ 16.7ms

## CSS3动画（简单动画的实现，如旋转等）

1. 依靠CSS3中提出的三个属性：transition、transform、animation
    1. transition：定义了元素在变化过程中是怎么样的，包含transition-property、transition-duration、transition-timing-function、transition-delay。
    1. transform：定义元素的变化结果，包含rotate、scale、skew、translate。
    1. animation：动画定义了动作的每一帧（@keyframes）有什么效果，包括animation-name，animation-duration、animation-timing-function、animation-delay、animation-iteration-count、animation-direction

## stylus/sass/less区别

1. 均具有“变量”、“混合”、“嵌套”、“继承”、“颜色混合”五大基本特性
1. Scss和LESS语法较为严谨，LESS要求一定要使用大括号“{}”，Scss和Stylus可以通过缩进表示层次与嵌套关系
1. Scss无全局变量的概念，LESS和Stylus有类似于其它语言的作用域概念
1. Sass是基于Ruby语言的，而LESS和Stylus可以基于NodeJS NPM下载相应库后进行编译；

## postcss的作用

1. 可以直观的理解为：它就是一个平台。为什么说它是一个平台呢？因为我们直接用它，感觉不能干什么事情，但是如果让一些插件在它上面跑，那么将会很强大
1. PostCSS 提供了一个解析器，它能够将 CSS 解析成抽象语法树
1. 通过在 PostCSS 这个平台上，我们能够开发一些插件，来处理我们的CSS，比如热门的：autoprefixer
1. postcss可以对sass处理过后的css再处理 最常见的就是autoprefixer

## 几种常见的CSS布局

1. 流体布局

```html
<div class="container">
    <div class="left"></div>
    <div class="right"></div>
    <div class="main"></div>
</div>
```

```css
.left {
    float: left;
    width: 100px;
    height: 200px;
    background: red;
}
.right {
    float: right;
    width: 200px;
    height: 200px;
    background: blue;
}
.main {
    margin-left: 120px;
    margin-right: 220px;
    height: 200px;
    background: green;
}
```

1. 圣杯布局

```html
<div class="container">
    <div class="main"></div>
    <div class="left"></div>
    <div class="right"></div>
</div>
```

```css
.container {
    margin-left: 120px;
    margin-right: 220px;
}
.main {
    float: left;
    width: 100%;
    height:300px;
    background: green;
}
.left {
    position: relative;
    left: -120px;
    float: left;
    height: 300px;
    width: 100px;
    margin-left: -100%;
    background: red;
}
.right {
    position: relative;
    right: -220px;
    float: right;
    height: 300px;
    width: 200px;
    margin-left: -200px;
    background: blue;
}
```

1. 双飞翼布局

```html
<div class="content">
    <div class="main"></div>
</div>
<div class="left"></div>
<div class="right"></div>
```

```css
.content {
    float: left;
    width: 100%;
}
.main {
    height: 200px;
    margin-left: 110px;
    margin-right: 220px;
    background: green;
}
.main::after {
    content: '';
    display: block;
    font-size:0;
    height: 0;
    zoom: 1;
    clear: both;
}
.left {
    float:left;
    height: 200px;
    width: 100px;
    margin-left: -100%;
    background: red;
}
.right {
    float: right;
    height: 200px;
    width: 200px;
    margin-left: -200px;
    background: blue;
}
```
