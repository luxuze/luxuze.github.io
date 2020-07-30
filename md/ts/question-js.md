# JavasSript

## 闭包

1. 闭包就是能够读取其他函数内部变量的函数
1. 闭包是指有权访问另一个函数作用域中变量的函数，创建闭包的最常见的方式就是在一个函数内创建另一个函数，通过另一个函数访问这个函数的局部变量,利用闭包可以突破作用链域
1. 闭包的特性：
    1. 函数内再嵌套函数
    1. 内部函数可以引用外层的参数和变量
    1. 参数和变量不会被垃圾回收机制回收
1. 使用闭包主要是为了设计私有的方法和变量。闭包的优点是可以避免全局变量的污染，缺点是闭包会常驻内存，会增大内存使用量，使用不当很容易造成内存泄露。在js中，函数即闭包，只有函数才会产生作用域的概念
1. 闭包 的最大用处有两个，一个是可以读取函数内部的变量，另一个就是让这些变量始终保持在内存中
1. 闭包的另一个用处，是封装对象的私有属性和私有方法
    1. 好处：能够实现封装和缓存等；
    1. 坏处：就是消耗内存、不正当使用会造成内存溢出的问题
1. 使用闭包的注意点
    1. 由于闭包会使得函数中的变量都被保存在内存中，内存消耗很大，所以不能滥用闭包，否则会造成网页的性能问题，在IE中可能导致内存泄露
    1. 解决方法是，在退出函数之前，将不使用的局部变量全部删除

## 作用域链

1. 作用域链的作用是保证执行环境里有权访问的变量和函数是有序的，作用域链的变量只能向上访问，变量访问到window对象即被终止，作用域链向下访问变量是不被允许的
简单的说，作用域就是变量与函数的可访问范围，即作用域控制着变量与函数的可见性和生命周期

## JavaScript原型，原型链 ? 有什么特点

1. 每个对象都会在其内部初始化一个属性，就是prototype(原型)，当我们访问一个对象的属性时
1. 如果这个对象内部不存在这个属性，那么他就会去prototype里找这个属性，这个prototype又会有自己的prototype，于是就这样一直找下去，也就是我们平时所说的原型链的概念
1. 关系：instance.constructor.prototype = instance.__proto__
1. 特点：
    1. JavaScript对象是通过引用来传递的，我们创建的每个新对象实体中并没有一份属于自己的原型副本。当我们修改原型时，与之相关的对象也会继承这一改变
    1. 当我们需要一个属性的时，Javascript引擎会先看当前对象中是否有这个属性， 如果没有的
    1. 就会查找他的Prototype对象是否有这个属性，如此递推下去，一直检索到 Object 内建对象

## 请解释什么是事件代理

1. 事件代理（Event Delegation），又称之为事件委托。是 JavaScript 中常用绑定事件的常用技巧。顾名思义，“事件代理”即是把原本需要绑定的事件委托给父元素，让父元素担当事件监听的职务。事件代理的原理是DOM元素的事件冒泡。使用事件代理的好处是可以提高性能
1. 可以大量节省内存占用，减少事件注册，比如在table上代理所有td的click事件就非常棒
1. 可以实现当新增子对象时无需再次对其绑定

## Javascript如何实现继承

1. 构造继承
1. 原型继承
1. 实例继承
1. 拷贝继承

原型prototype机制或apply和call方法去实现较简单，建议使用构造函数与原型混合方式

```javascript
function Parent(){
        this.name = 'wang';
    }

    function Child(){
        this.age = 28;
    }
    Child.prototype = new Parent();//继承了Parent，通过原型

    var demo = new Child();
    alert(demo.age);
    alert(demo.name);//得到被继承的属性
}
```

## This对象的理解

1. this总是指向函数的直接调用者（而非间接调用者）
1. 如果有new关键字，this指向new出来的那个对象
1. 在事件中，this指向触发这个事件的对象，特殊的是，IE中的attachEvent中的this总是指向全局对象Window

## 事件模型

1. W3C中定义事件的发生经历三个阶段：捕获阶段（capturing）、目标阶段（targetin）、冒泡阶段（bubbling）
1. 冒泡型事件：当你使用事件冒泡时，子级元素先触发，父级元素后触发
1. 捕获型事件：当你使用事件捕获时，父级元素先触发，子级元素后触发
1. DOM事件流：同时支持两种事件模型：捕获型事件和冒泡型事件
1. 阻止冒泡：在W3c中，使用stopPropagation（）方法；在IE下设置cancelBubble = true
1. 阻止捕获：阻止事件的默认行为，例如click - <a>后的跳转。在W3c中，使用preventDefault（）方法，在IE下设置window.event.returnValue = false

## new操作符具体干了什么呢

1. 创建一个空对象，并且 this 变量引用该对象，同时还继承了该函数的原型
1. 属性和方法被加入到 this 引用的对象中
1. 新创建的对象由 this 所引用，并且最后隐式的返回 this

## Ajax原理

1. Ajax的原理简单来说是在用户和服务器之间加了—个中间层(AJAX引擎)，通过XmlHttpRequest对象来向服务器发异步请求，从服务器获得数据，然后用javascript来操作DOM而更新页面。使用户操作与服务器响应异步化。这其中最关键的一步就是从服务器获得请求数据
1. Ajax的过程只涉及JavaScript、XMLHttpRequest和DOM。XMLHttpRequest是ajax的核心机制

```javascript
 // 1. 创建连接
    var xhr = null;
    xhr = new XMLHttpRequest()
    // 2. 连接服务器
    xhr.open('get', url, true)
    // 3. 发送请求
    xhr.send(null);
    // 4. 接受请求
    xhr.onreadystatechange = function(){
        if(xhr.readyState == 4){
            if(xhr.status == 200){
                success(xhr.responseText);
            } else { // fail
                fail && fail(xhr.status);
            }
        }
    }
```

## Ajax优缺点

1. 优点：
    1. 通过异步模式，提升了用户体验.
    1. 优化了浏览器和服务器之间的传输，减少不必要的数据往返，减少了带宽占用.
    1. Ajax在客户端运行，承担了一部分本来由服务器承担的工作，减少了大用户量下的服务器负载。
    1. Ajax可以实现动态不刷新（局部刷新）
1. 缺点：
    1. 安全问题 AJAX暴露了与服务器交互的细节。
    1. 对搜索引擎的支持比较弱。
    1. 不容易调试。

## 如何解决跨域问题

1. jsonp、 iframe、window.name、window.postMessage、服务器上设置代理页面

## 模块化开发怎么做

1. 立即执行函数,不暴露私有成员

```javascript
var module1 = (function(){
　　　　var _count = 0;
　　　　var m1 = function(){
　　　　　　//...
　　　　};
　　　　var m2 = function(){
　　　　　　//...
　　　　};
　　　　return {
　　　　　　m1 : m1,
　　　　　　m2 : m2
　　　　};
　　})();
```

## 那些操作会造成内存泄漏

1. 内存泄漏指任何对象在您不再拥有或需要它之后仍然存在
1. setTimeout 的第一个参数使用字符串而非函数的话，会引发内存泄漏
1. 闭包使用不当

## 对webpack的看法

1. WebPack 是一个模块打包工具，你可以使用WebPack管理你的模块依赖，并编绎输出模块们所需的静态文件。它能够很好地管理、打包Web开发中所用到的HTML、Javascript、CSS以及各种静态文件（图片、字体等），让开发过程更加高效。对于不同类型的资源
1. webpack有对应的模块加载器。webpack模块打包器会分析模块间的依赖关系，最后 生成了优化且合并后的静态资源

## 对AMD和Commonjs的理解

1. CommonJS是服务器端模块的规范，Node.js采用了这个规范。CommonJS规范加载模块是同步的，也就是说，只有加载完成，才能执行后面的操作。AMD规范则是非同步加载模块，允许指定回调函数
1. AMD推荐的风格通过返回一个对象做为模块对象，CommonJS的风格通过对module.exports或exports的属性赋值来达到暴露模块对象的目的

## 常见web安全及防护原理

1. sql注入,通过把SQL命令插入到Web表单递交或输入域名或页面请求的查询字符串，最终达到欺骗服务器执行恶意的SQL命令
1. 总的来说有以下几点
    1. 永远不要信任用户的输入，要对用户的输入进行校验，可以通过正则表达式，或限制长度，对单引号和双"-"进行转换等
    1. 永远不要使用动态拼装SQL，可以使用参数化的SQL或者直接使用存储过程进行数据查询存取
    1. 永远不要使用管理员权限的数据库连接，为每个应用使用单独的权限有限的数据库连接
    1. 不要把机密信息明文存放，请加密或者hash掉密码和敏感的信息
1. XSS原理及防范
    1. Xss(cross-site scripting)攻击指的是攻击者往Web页面里插入恶意html标签或者javascript代码。比如：攻击者在论坛中放一个看似安全的链接，骗取用户点击后，窃取cookie中的用户私密信息；或者攻击者在论坛中加一个恶意表单，当用户提交表单的时候，1. 却把信息传送到攻击者的服务器中，而不是用户原本以为的信任站点
    1. XSS防范方法
    1. 首先代码里对用户输入的地方和变量都需要仔细检查长度和对”<”,”>”,”;”,”’”等字符做过滤；其次任何内容写到页面之前都必须加以encode，避免不小心把html tag 弄出来。这一个层面做好，至少可以堵住超过一半的XSS 攻击
1. XSS与CSRF有什么区别吗？
    1. XSS是获取信息，不需要提前知道其他用户页面的代码和数据包。CSRF是代替用户完成指定的动作，需要知道其他用户页面的代码和数据包。要完成一次CSRF攻击，受害者必须依次完成两个步骤
    1. 登录受信任网站A，并在本地生成Cookie
    1. 在不登出A的情况下，访问危险网站B
1. CSRF的防御
    1. 服务端的CSRF方式方法很多样，但总的思想都是一致的，就是在客户端页面增加伪随机数
    1. 通过验证码的方法

## 为什么要有同源限制

1. 同源策略指的是：协议，域名，端口相同，同源策略是一种安全协议
1. 举例说明：比如一个黑客程序，他利用Iframe把真正的银行登录页面嵌到他的页面上，当你使用真实的用户名，密码登录时，他的页面就可以通过Javascript读取到你的表单中input中的内容，这样用户名，密码就轻松到手了。

## offsetWidth/offsetHeight,clientWidth/clientHeight与scrollWidth/scrollHeight的区别

1. offsetWidth/offsetHeight返回值包含content + padding + border，效果与e.getBoundingClientRect()相同
1. clientWidth/clientHeight返回值只包含content + padding，如果有滚动条，也不包含滚动条
1. scrollWidth/scrollHeight返回值包含content + padding + 溢出内容的尺寸

## javascript有哪些方法定义对象

1. 对象字面量： var obj = {};
1. 构造函数： var obj = new Object();
1. Object.create(): var obj = Object.create(Object.prototype);

## 常见兼容性问题

1. png24位的图片在iE6浏览器上出现背景，解决方案是做成PNG8
1. 浏览器默认的margin和padding不同。解决方案是加一个全局的*{margin:0;padding:0;}来统一,，但是全局效率很低，一般是如下这样解决：

```css
body,ul,li,ol,dl,dt,dd,form,input,h1,h2,h3,h4,h5,h6,p{
    margin:0;
    padding:0;
}
```

1. IE下,event对象有x,y属性,但是没有pageX,pageY属性
1. Firefox下,event对象有pageX,pageY属性,但是没有x,y属性.

## 说说你对promise的了解

1. 依照 Promise/A+ 的定义，Promise 有四种状态：
1. pending: 初始状态, 非 fulfilled 或 rejected.
1. fulfilled: 成功的操作.
1. rejected: 失败的操作.
1. settled: Promise已被fulfilled或rejected，且不是pending
1. 另外， fulfilled与 rejected一起合称 settled
1. Promise 对象用来进行延迟(deferred) 和异步(asynchronous) 计算

## Promise 的构造函数

1. 构造一个 Promise，最基本的用法如下：

```javascript
var promise = new Promise(function(resolve, reject) {
        if (...) {  // succeed
            resolve(result);
        } else {   // fails
            reject(Error(errMessage));
        }
    });
```

1. Promise 实例拥有 then 方法（具有 then 方法的对象，通常被称为thenable）。它的使用方法如下：
1. promise.then(onFulfilled, onRejected)
1. 接收两个函数作为参数，一个在 fulfilled 的时候被调用，一个在rejected的时候被调用，接收参数就是 future，onFulfilled 对应resolve, onRejected对应 reject

## vue、react、angular

1. Vue.js 一个用于创建 web 交互界面的库，是一个精简的 MVVM。它通过双向数据绑定把 View 层和 Model层连接了起来。实际的 DOM 封装和输出格式都被抽象为了Directives 和 Filters
1. AngularJS 是一个比较完善的前端MVVM框架，包含模板，数据双向绑定，路由，模块化，服务，依赖注入等所有功能，模板功能强大丰富，自带了丰富的 Angular指令
1. React 仅仅是 VIEW 层是facebook公司。推出的一个用于构建UI的一个库，能够实现服务器端的渲染。用了virtual dom，所以性能很好。

1. ## 介绍js有哪些内置对象

1. Object 是 JavaScript 中所有对象的父对象
1. 数据封装类对象：Object、Array、Boolean、Number 和 String
1. 其他对象：Function、Arguments、Math、Date、RegExp、Error

## JavaScript有几种类型的值？，你能画一下他们的内存图吗

1. 栈：原始数据类型（Undefined，Null，Boolean，Number、String）
1. 堆：引用数据类型（对象、数组和函数）
1. 两种类型的区别是：存储位置不同；
    1. 原始数据类型直接存储在栈(stack)中的简单数据段，占据空间小、大小固定，属于被频繁使用数据，所以放入栈中存储；
    1. 引用数据类型存储在堆(heap)中的对象,占据空间大、大小不固定,如果存储在栈中，将会影响程序运行的性能；引用数据类型在栈中存储了指针，该指针指向堆中该实体的起始地址。当解释器寻找引用值时，会首先检索其
    1. 在栈中的地址，取得地址后从堆中获得实体

## eval是做什么的

1. 它的功能是把对应的字符串解析成JS代码并运行
1. 应该避免使用eval，不安全，非常耗性能（2次，一次解析成js语句，一次执行）
1. 由JSON字符串转换为JSON对象的时候可以用eval，var obj =eval('('+ str +')')

## 渐进增强和优雅降级

1. 渐进增强 ：针对低版本浏览器进行构建页面，保证最基本的功能，然后再针对高级浏览器进行效果、交互等改进和追加功能达到更好的用户体验。
1. 优雅降级 ：一开始就构建完整的功能，然后再针对低版本浏览器进行兼容

## attribute和property的区别是什么

1. attribute是dom元素在文档中作为html标签拥有的属性；
1. property就是dom元素在js中作为对象拥有的属性。
1. 对于html的标准属性来说，attribute和property是同步的，是会自动更新的
1. 但是对于自定义的属性来说，他们是不同步的

## 如何通过JS判断一个数组

1. instanceof方法: instanceof 运算符是用来测试一个对象是否在其原型链原型构造函数的属性

```javascript
var arr = [];
arr instanceof Array; // true
```

1. constructor方法:constructor属性返回对创建此对象的数组函数的引用，就是返回对象相对应的构造函数

```javascript
var arr = [];
arr.constructor == Array; //true
```

1. 最简单的方法, 这种写法，是 jQuery 正在使用的

```javascript
Object.prototype.toString.call(value) == '[object Array]'
// 利用这个方法，可以写一个返回数据类型的方法
var isType = function (obj) {
     return Object.prototype.toString.call(obj).slice(8,-1);
}
```

1. ES5新增方法isArray()

```javascript
var a = new Array(123);
var b = new Date();
console.log(Array.isArray(a)); //true
console.log(Array.isArray(b)); //false
```

## map与forEach的区别

1. forEach方法，是最基本的方法，就是遍历与循环，默认有3个传参：分别是遍历的数组内容item、数组索引index、和当前遍历数组Array
1. map方法，基本用法与forEach一致，但是不同的，它会返回一个新的数组，所以在callback需要有return值，如果没有，会返回undefined

## Vue的双向绑定数据的原理

1. vue.js 则是采用数据劫持结合发布者-订阅者模式的方式，通过Object.defineProperty()来劫持各个属性的setter，getter，在数据变动时发布消息给订阅者，触发相应的监听回调

## TypeScript Singleton 单例模式

[Stack Over Flow Link](https://stackoverflow.com/questions/30174078/how-to-define-singleton-in-typescript)

```typescript
class Singleton {
  //Assign "new Singleton()" here to avoid lazy initialisation
  private static instance: Singleton;

  constructor(sp:{}) {
    if (Singleton.instance) return Singleton.instance;
    this.someprops = sp;
    Singleton.instance = this;
  }
  someprops:{};
}
```

## typescript string/String

1. [Typescript: difference between String and string](https://stackoverflow.com/questions/14727044/typescript-difference-between-string-and-string)
