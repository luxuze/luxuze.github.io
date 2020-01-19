# TypeScript Singleton 单例模式

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
