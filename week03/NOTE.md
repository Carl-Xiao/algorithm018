# 心得

这周东西确实花时间相对较少，又遇到了这种递归回溯法。发现与解决二叉树感觉如果进入了新的世界,在国际站上找到了一些比较优秀的解法。参考才能完成当前的作业


加油把~~day day up



## 牢记模板

```java
public void recur(int level, int param) { 
  // terminator 
  if (level > MAX_LEVEL) { 
    // process result 
    return; 
  }
  // process current logic 
  process(level, param); 
  // drill down 
  recur( level: level + 1, newParam); 
  // restore current status 
 
}
```
- 找到当前层的边界值
- 处理当前层的逻辑
- 接着向下走
```golang

func recur(level int ,parm int){
	 // terminator 
	 if level > MAX_LEVEL { 
		// process result 
		return; 
	  }
	  // process current logic 
	  process(level, param); 
	  // drill down 
	  recur(level + 1, newParam); 
	  // restore current status 
}
```