## 工厂模式
> 工厂模式（Factory Pattern）是 Java 中最常用的设计模式之一。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。
  在工厂模式中，我们在创建对象时不会对客户端暴露创建逻辑，并且是通过使用一个共同的接口来指向新创建的对象。
  
工厂模式不适合go里面使用，需要哪个实例，建议直接从实例去New出来，基于以下原因：
 
 - 工厂模式需要传入实例key string，这个key 需要多余的维护，并且容易写错。这个key一旦写错。
 无法享受go的编译前错误，将他报出来，而只有运行期才可以报错。
 - 和通过pkg.Instance的调用相比，没有管理上，性能上的优势，仅仅是为了模仿java而衍生出来的东西。
 - 继承了一些接口的缺点，如果一个功能，可以抛开接口来实现，那么就不要使用接口。接口在编程逻辑学里很美，但是在实际开发中有以下问题：
    1. 无法通过调用的位置，点出实例的实现位置。而会跳出接口声明，进而凭感觉去挑选哪个实例是你要看的。
    2. 业务侧容易做需求迭代。一个接口很难在声明初期就预设出他的所有限制要求，而任何的变动，都依赖所有实例做变动，他跨越了作者、时间、版本。
    