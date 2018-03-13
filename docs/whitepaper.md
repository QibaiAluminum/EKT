# <center>EKT白皮书</center>
### 目录
<a href="#cp1">1、EKT公链简介</a><br/>
<a href="#cp2">2、多链多共识</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp21">2.1 多链模型</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp22">2.2 EKT为什么要设计多链模型</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp23">2.3 EKT要怎么实现多条链互相隔离又使用同一套地址呢</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp24">2.4 开发者为何不直接fork EKT主链的代码部署自己的主链,而要在EKT上发布自己的侧链</a><br/>
<a href="#cp3">3、跨公链</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp31">3.1 EKT公链中的跨链资产转移</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp32">3.2 跨公链交易</a><br/>
<a href="#cp4">4、DAPP</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp41">4.1 Paxos</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp42">4.2 智能合约语言</a><br/>
<a href="#cp5">5、EKT适合哪些开发者</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp51">5.1 刚刚接触区块链的传统互联网开发者</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp52">5.2 对 DAPP感兴趣,想编写一个自己的 DAPP的开发者</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp53">5.3 想发行自己的 token,但是其他数据和逻辑不上链的开发者</a><br/>
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<a href="#cp54">5.4	对于区块链有深入了解的开发者</a><br/>


<h1 id="cp1">1、EKT公链简介</h1>
<p>比特币是第一代的去中心化的货币,实现了一个完整的去中心化记账的功能,但是开发者无法针对比特币开发去中心化的应用,于是就有了第二代区块链项目以太坊,以太坊是一个可编程平台,开发者可以开发出智能合约提交到以太坊的节点上,这样一来开发者就可以发布自己的去中心化应用程序,ERC20代币就是智能合约一个很好的应用，但是在以太坊上开发其他类型的应用的难度是非常大的。非常出名的加密猫就是一个在以太坊上运行的应用，但是有过程序设计经验的同学可以看出来，其实以太坊的加密猫在逻辑上是非常简单的，而且非常适合以太坊提供的键值对的数据模型，如果我们开发一个业务逻辑相对复杂的应用的话，仍然是没有什么特别好的选择。EKT 公链设计之初就是要实现可以跨公链进行转账的协议,同时提供功能更加完善的智能合约虚拟机,针对开发者开发不同类型的DAPP还会有不同的编程接口,所以开发者可以很方便在EKT上发行自己的币或者实现自己其他类型的DAPP,同时还可以很容易调用跨链操作API,这些都会体现在智能合约的内置接口上。</p>
<h1 id="cp2">2、多链多共识</h1>
<h2 id="cp21">2.1 多链模型</h2>
在EKT公链中，会存在多条并行的链，每条链上都会有一个主币（即交易费需要用的币）和很多普通币，多条链共享同一套用户地址（值得一提的是，在EKT中，用户的地址不是通过公钥计算出来的，而是通过申请得到的，用户可以更换自己的地址对应的公私钥对），不同的链可以存在不同的共识机制，默认是DPOS，开发者可以实现Consensus接口实现自己的接口部署自己的节点，另外我们后面也会加入对POW和POS共识的支持，相信这对于一些开发能力较弱的团队或者对区块链没有太多深入了解的团队上链开发DAPP是一个福音。
<h2 id="cp22">2.2 EKT为什么要设计多链模型</h2>
我们都知道在比特币中，TPS一般都是在7左右，以太坊15TPS，在一个去中心化的网络里，为了提升安全性是需要牺牲一定的性能的。相信很多开发者都听说过CAP理论，即consistency一致性、availability可用性，partition-tolerance分区容错性，在数据库系统中，关系型数据库选择了CA，NoSql选择了AP，三者不可兼得。那么在区块链中是一种什么情况呢？首先区块链是建立在一个不可信的网络节点上的，所以所有数据备份在同一时间节点上是不可能给具有相同的值的，所以在区块链上首先放弃了C，但是并不是在区块链上的数据并不是不一致的，在比特币中，为了保证大家最终的数据一致性，引入了最长链的算法，即包含区块数量更多的链是主链，所有的区块去同步最长链。由此看来区块链中，CAP三者也是不可兼得的，那么是否得到剩下的两个呢？我们都知道区块链是基于P2P网络的是一个服务，所有节点都可以提供数据，所以可用性是一致都有的，那么分区容错性通过最长链的同步算法就可以实现，所以区块链实现了AP优先和最终一致性。那么为什么一般情况下选择AP的NoSql可以达到很高的TPS而区块链不行呢？原因很简单，TPS的大小取决于一个公式：区块内Transaction的数量/区块产生的时间，也就是说，如果想要达到很高的TPS就必须要提高区块大小或者缩短区块产生的时间。我们都知道在比特币中，采用的是UTXO模型，使用Merkle树存储交易信息，所以树的深度就是log(n)，如果要增加区块大小，消耗的磁盘的速度将会非常快，因此不是一个明智的选择。而减少区块产生的时间看起来是一个比较合适的方案，但是如果降低挖矿难度的话，在节点数量非常多的情况下，最长链的算法并不能保证不同节点的最终一致性，所以在POW共识下是非常不合适的。那么如果才能提高TPS呢？相信答案现在已经很清楚了，那就是并行运行的多条链，对实时性要求比较高的采用DPOS共识或者瑞波共识，对去中心化程度要求比较高的可以使用POW共识，多条链互相隔离，就可以实现更高的TPS了，同时还满足了对不同去中心化程度的应用的需求。
<h2 id="cp23">2.3 EKT要怎么实现多条链互相隔离又使用同一套地址呢</h2>
相信这个问题是对很多开发者来说最难理解的了，因为我们上面提到过，在EKT中用户的地址不是通过公钥计算出来的，而是通过申请得到的，用户可以更换自己的地址对应的公私钥对，那么如果用户更换了自己的公私钥，其他链是如何校验用户的签名信息呢？方法有两种：一种是对安全性要求相对不那么高的方法，就是同步主链区块，使用最新区块校验用户的公钥信息，即使同步主链区块出现延迟也是没有关系的，因为用户修改自己的公私钥对并不代表自己的私钥已经被破解了，2的138次方的复杂度还是很大的，所以即使出现1-2个区块的延迟，旧签名仍是有效的。那么如果用户更新了自己的私钥在使用新私钥对消息进行签名的时候，侧链校验失败了怎么办？这时候在区块上可以使用第二种方法：RPC调用主链的校验方法，即将地址、消息和签名发给主链的多个节点，节点处理完成之后返回签名校验结果，一般情况下超过n/2+1个委托人节点签名通过即可以视为签名通过，可以证明当前用户拥有对此地址的权限。
<h2 id="cp24">2.4 开发者为何不直接fork EKT主链的代码部署自己的主链，而要在EKT上发布自己的侧链</h2>
能想到这个问题的读者相信都是有一定编码能力或者深谙编码原理的老司机了，因为发布自己的主链然后新增一些数据结构就可以完成一个去中心化应用的开发。EKT之所以有信心让大家在EKT上发布自己的侧链而不是fork代码有以下几个原因：<br/>
<h3 id="cp241">2.4.1 与EKT链上的其他链共享用户</h3>
在这个世界上，每个人都应该有且只有一个自己的id，并且是保存在区块链上的，如果大家对用户地址、签名拥有共识，我相信没有开发者会想着在给每个用户“取个名字”了。在EKT公链中，提供了两种不同安全性但都非常高效的用户机制，这是一个能让所有开发者达成共识并且受益的方式，我相信大家都会选择最好的方法。
<h3 id="cp242">2.4.2 与EKT上活跃的开发者共享开发库</h3>
EKT公链拥有很多的开发者，并且拥有非常大的开发者社区，与其自己闭门造轮子不如大家共享自己造出来的轮子，从github就可以看出来开源社区对programer's world影响有多大，在EKT社区上共享开发资源何尝不是一种共识呢？
<h3 id="cp243">2.4.3 EKT主链与侧链并没有主次之分</h3>
在EKT公链上，并没有主次之分，所有链都是公平的，除了用户公私钥只能放在主链上去处理之外，其他功能都是相同的。比如对于一个侧链的owner来说，在此侧链上的手续费是自己的主币，而不是EKT，所以在EKT上发布自己的侧链是一个非常好的选择。
<h1 id="cp3">3、跨公链</h1>
在EKT上，主链和侧链之间或者侧链和侧链之间除了资产打包转移之外是不存在所谓的“跨链”操作，同时也支持跨公链操作，总结起来就是：多链多共识，一链一主币，原生支持公链内跨链操作，支持跨公链操作。那么我先说一下EKT中的跨链资产转移。
<h2 id="cp31">3.1 EKT公链中的跨链资产转移</h2>
在EKT公链中，是不存在所谓的跨链交易的，用户A和用户B分别拥有Token T1和Token T2，T1和T2在不同的链上，我们将操作`A将拥有的T1给B`定义为Tx1，将操作`B将拥有的T2给A`定义为Tx2，这个时候Tx1是在Token T1所在的链上进行的，Tx2是在Token T2所在的链上进行的，手续费分别为两个链上的主币，都是普通的转账操作，因为上面提到过不同的链共享了同一套用户地址，所以并没有产生跨链操作，也就是说在EKT中是原生支持“跨链操作”的。<br/>
那既然EKT公链中原生支持跨链操作，跨链资产转移又是什么呢？跨链资产转移是指更改当前token打包的链，即一个token只会在一条链上进行打包，在EKT公链上的跨链资产转移可能有一下几个原因：当前网络比较拥堵（比如当前token所在的链使用的是POW共识，难度设置的过高）、想更换自己的共识协议、当前区块的交易费用过高、想自己创建一条侧链成为主币。进行跨链资产转移需要指定转移的高度和目标链，资产转移之后也就遵循新链的共识机制。
<h2 id="cp32">3.2 跨公链交易</h2>
在EKT中是支持对外的跨公链交易的，但是是需要一定的步骤的。
<h3 id="cp321">3.2.1 跨公链的用户共识</h3>
两个公链如果要进行跨链的资产交易的话，首先要对用户拥有共识，即从一个公链的一个地址转移到另一个链的地址需要双方都对对方的公链进行注册，最简单有效的方法就是`长地址`，在EKT中，存在两种地址：`内部地址`和`外部地址`，内部地址就是在EKT公链中各个链使用的地址，用于在EKT主链和侧链之间转账和其他DAPP的开发使用。同时EKT中还存在外部地址，外部地址拥有68byte，前4byte存储外部地址的公链id的长度和在其公链上内部地址的长度，后n位存储的是公链id和内部地址，中间用0x00填充。
<h3 id="cp322">3.2.2 跨公链的token共识</h3>
与用户地址相同，token的地址也是分为`内部地址`和`外部地址`，与用户地址相同，token的内部地址就是在EKT公链中各个链使用的地址，用于在EKT主链和侧链之间转账和其他DAPP的开发使用。同时EKT中Token也存在外部地址，同样拥有68byte，前4byte存储外部地址的公链id的长度和在其公链上内部token地址的长度，后n位存储的是公链id和内部地址，中间用0x00填充。
<h3 id="cp323">3.2.3 跨公链的交易共识</h3>
在EKT的跨公链操作中，对跨链协议是这样定义的：假如现在存在两条公链EKT定义为`PC1`和ETH定义为`PC2`，在`PC1`上有一个用户`user1`，地址位`address1`， 在`PC2`上拥有一个用户`user2`，在`PC1`上有一个Token定义为`T1`。如果`user1`要将自己在`PC1`上拥有的`T1`转移给`PC2`上的`user2`，那么在`PC2`上必须已经对`PC1`这条公链和`T1`进行了注册，这时就有了第一个报文：<a href="#handshake">跨链握手报文</a> ，当`PC2`对上述报文回复正确之后，`user1`才可以将自己的`T1`转移到了`PC2`。在此跨公链转账打包成功之后，`PC1`需要将交易信息和区块头信息发送给`PC2`进行校验，这就是第二个报文：<a href="#transaction">跨链交易报文</a> ，`PC2`校验之后存入自己的区块链中之后需要给`PC1`回复第三个报文：<a href="#confirm">跨链确认报文</a>，否则`PC1`会继续给`PC2`发送<a href="#transaction">跨链交易报文</a>，直到自己收到`PC2`的第三个报文：<a href="#confirm">跨链确认报文</a>。给`user1`将自己的`T1`转移到了`PC2`上之后，这些Token在`PC2`上怎么转移对于`PC1`来说是无关紧要的，因为在`PC1`上记录的是将`T1`转移到了`PC2`，而在`PC2`向`PC1`发送<a href="#transaction">跨链交易报文</a>的时候，`PC1`也是不会校验在`PC2`上每个地址持有`T1`的数量，而是校验在`PC1`链中当前的`PC2`公链id拥有的`T1`的数量。
<h3 id="cp324">3.2.4 跨公链报文协议</h3>
EKT在定义报文协议的时候为了兼容大多数公链，使用了HTTP协议作为基础的报文协议，请求和响应的Content-Type必须为application/json，请求参数和响应参数必须与规定相同。详见 [跨链报文协议](./crosschain.md)
<h4 id="cp3241">3.2.4.1跨链握手协议</h4>
    详见 [跨链报文协议](./crosschain.md)
<h4 id="cp3242">3.2.4.2跨链交易报文</h4>
    详见 [跨链报文协议](./crosschain.md)
<h4 id="cp3243">3.2.4.3跨链确认报文</h4>
    详见 [跨链报文协议](./crosschain.md)
<h1 id="cp4">4、DAPP-无区块</h1>
EKT的DAPP分为DAPP服务端和DAPP客户端，DAPP服务端就是智能合约，DAPP客户端是基于JavaScript的一个框架，用户可以根据这个框架实现自己的客户端功能。<br/>
EKT的DAPP有啥优势呢？<br/>
1、上手简单，因为EKT在设计智能合约的时候考虑过现有智能合约的缺陷，上手难度大，与业务无关操作比较多，而且目前大部分的智能合约都是基于KV进行存储的，没有对上层智能合约提供可靠易用的SDK，门槛很高。
2、基于DPOS+PAXOS的智能合约共识机制，可以让DAPP开发者不必关注共识算法，专注于自己业务逻辑的开发。我们都知道Paxos在跨区域（跨机房）环境下几乎不可用，EKT使用了什么方法让可以在智能合约的共识上加入了Paxos呢？详见 <a href="cp412">DPOS+Paxos</a>
<h2 id="cp41">4.1 Paxos</h2>
什么是Paxos，相信很多对分布式一致性的有过研究的都会对Paxos有很深的认识，Google Chubby的作者Mike Burrows就曾说过这个世界上只有一种一致性算法，那就是Paxos，其它的算法都是残次品。在Paxos中有一句话是这样的：如果各节点的初始状态一致，每个节点执行相同的操作序列，那么他们最后能得到一个一致的状态。相信对于这句话大家都是认可的，但是如果实现在各个节点执行相同的操作序列就是一个难题了。在Paxos中，有三种角色，Proposer、Acceptor和Learner，同一个进程可能扮演多个不同的角色。从宏观上看，比特币的一个区块就是一个带有顺序的操作序列的集合，挖矿过程中就是大家在争夺Proposer角色的过程，当有一个Proposer发起一个提案之后，那么其他节点都会成为Acceptor，当大家对当前提案有了一个共识之后，所有节点都会成为Learner，去同步区块，然后争夺下一个节点的Proposer的权限。
<h3 id="cp411">4.1.1 什么是Paxos</h3>
Paxos的介绍可以说是非常多了，在这里我就不做过多解释了，给大家推荐一个博客，写的深入浅出。 <a href="http://www.cnblogs.com/linbingdong/p/6253479.html">分布式系列文章——Paxos算法原理与推导</a>
<h3 id="cp412">4.1.2 DPOS+Paxos？</h3>
我们都知道DPOS是类似于“人大代表”制度，选出一定数量的节点来进行存储和计算，那么DPOS和Paxos又能擦出什么样的火花呢？重复依据上面的话：如果各节点的初始状态一致，每个节点执行相同的操作序列，那么他们最后能得到一个一致的状态。那么其实在智能合约中我们要做的就是让消息按照一定的顺序执行了。在EKT公链中，有两种粒度的“锁”，第一种锁是用户的“锁”，即每个用户对于每个智能合约都存在一个nonce值，每次调用API当前nonce自增1，实现用户API的自然锁。那么如果一个用户对A节点发送的nonce值为`nonce1`的API对应的是`API1`，对B节点发送的nonce值为`nonce1`的API对应的是`API2`，类似于“双花问题”，那么又该如何处理呢？在EKT的设计中，每个节点都是Proposer，也都是Acceptor，当节点A收到一个客户端发送的消息`API1`之后，会对nonce进行判断，如果当前用户的nonce与数据库中刚好差1，那么当前节点会发起一个提案，如果其他节点同意了当前提案，则用户当前的nonce加一，这个时候节点B收到了`API2`，B判断当前的用户的nonce与当前nonce相等，则放弃当前API。所以对于用户的“双花攻击”，DPOS+Paxos是自带防御的。那么如果DAPP开发这有一个这样的功能：对外提供100个名额，先到先得，此时EKT公链中提供的第一种“用户锁”看起来就没什么用了，不过没关系，在EKT中提供了另一种锁：“全局时间锁”。全局时间锁是EKT公链的最大的一个创新，当Acceptor对一个提案进行签名的时候会带上当前的时间戳，当一半以上的Acceptor对一个提案进行签名之后，所有节点都需要执行这个提案，假如节点数为n，那么前n/2+1个同意节点的签名时间戳的平均值会被带入智能合约，实现全局时间。我们可以看出，实现了全局时间之后，很多问题都可以迎刃而解：资源竞争问题、客户端时间不可信问题、不同节点时间不同等等。
<h3 id="cp413">4.1.3 为何在智能合约上使用Paxos算法</h3>
继续重复一下上面的一句话：如果各节点的初始状态一致，每个节点执行相同的操作序列，那么他们最后能得到一个一致的状态。在这里根据Paxos的理论表达一个观点：在节点间延迟足够小的情况下，转账功能使用有向无环图的方式才是最优方式。在DPOS的情况下，是可以使用DAG代替区块的概念的，使用的人越多，理论上的TPS越高，但是在POW和POS的共识情况下，容易出现用户向两个网络延迟比较大的节点发送两个不同的交易，双花导致分叉，DAG的同步算法又无法足够成熟能支持最长深度优先，所以为了实现多链多共识的伟大愿景EKT在Token交易上放弃了Paxos算法和DAG模型，使用了区块概念。但是在智能合约上对于去中心化成都要求没那么高，所以EKT的智能合约采用了DPOS+Paxos共识，与客户端的API可以实现毫秒级确认和所有节点的数据同步，所以在EKT的智能合约上只支持DPOS+Paxos共识。
<h2 id="cp42">4.2 智能合约语言</h2>
EKT的新语言名叫AWM，是一个事件驱动的语言。事件分两种类型：用户事件和系统事件。用户事件是指DAPP客户端的调用，经过DPOS+Paxos的共识之后会发送一个事件到智能合约中。系统事件是指区块完成打包、某个txId处理完成或者智能合约订阅的一些其他事件。另外AWM支持面向对象，开发者可以使用面向对象的一些思想来编写代码。
<h3 id="cp421">4.2.1 语言特性</h3>
AWM拥有三个特性：<br/>
1、事件驱动。AWM是一个事件驱动的语言，也就是与一般语言（如C/C++，Java， Golang）不同，找不到main函数，所有的代码是围绕着事件去执行的。
2、面向对象。AWM是支持面向对象的，而且是一个天然模块化设计的一个语言，开发者可以把一些常用的结构体封装成一个对象，对对象封装一些方法，在事件驱动的函数里对对象的成员变量和成员方法进行调用。
3、模块化设计。在AWM中，允许把代码放在不同的package里面，我们鼓励用户把不同模块的代码放在不同的package中，供上层调用。
<h3 id="cp422">4.2.2 AWM VM</h3>
AWM VM是一个AWM的执行环境，相信很多人都在抱怨solidity的难以调试，相信AWM VM的到来会给大家一个惊喜，虽然在AWM中没有main的存在，但是用户可以在本地运行AWM VM，另外AWM VM还会对外提供API可以让开发者模拟事件，用户可以很轻易的使用Python、Shell或者JavaScript等语言编写测试案例，用来测试自己的智能合约是否能满足所有的测试案例，相信TDD是可以让大家极大的提高DAPP开发效率的。
<h3 id="cp423">4.2.3 原生函数</h3>
在AWM的底层是有许多基础库可以调用的，而且基础库是会不断完善的。值得一提的是原生数据库操作函数，在EKT中，智能合约是可以直接操作数据库的，当然是在封装在基础的磁盘操作之上的，我们会为用户提供三种方式的数据库操作：SQL、文档型（类似MongoDB）和KV型，开发者只需要导入响应的包即可完成响应的操作，相信SQL和文档数据库对大家开发速度和难度的提升是相当大的吧。
<h1 id="cp5">5、EKT适合哪些开发者</h1>
<h2 id="cp51">5.1 刚刚接触区块链的传统互联网开发者</h2>
EKT的模块化设计和完善的注释文档相信对于许多互联网开发者是非常友好的，一个传统的互联网开发者可以不用管共识和P2P就可以编写自己的DAPP，还可以对EKT公链进行学习，部署自己的私有链。
<h2 id="cp52">5.2 对DAPP感兴趣，想编写一个自己的DAPP的开发者</h2>
AWM对于DAPP开发的支持可以说是非常完善了，一个DAPP的开发者只需要一个小时的时间就可以动手写自己的DAPP了，另外还可以寻求EKT官方帮助，实现自己的DAPP。
<h2 id="cp53">5.3 想发行自己的token，但是其他数据和逻辑不上链的开发者</h2>
EKT支持但不鼓励这种类型的应用。
<h2 id="cp54">5.4 对于区块链有深入了解的开发者</h2>
这类开发者可以想EKT公链发PR，上传自己的共识或者其他模块的代码，然后可以发布自己的主链，共享EKT和其他链的用户基础，也可以发布一些自己写的库，与EKT团队一起完善国内最先进的公链。