# KPI-HR-GITLAB
### VERSION = v1.0.5

> + ## 基本功能：
>> 1. 通过与自建gitlab中代码管理系统交互，获取每个人的代码提交情况以及分析；
>> 2. 获取gitlab中所有人员的信息，包括：ID；姓名；用户名；头像地址；email；是否活跃；
>> 3. 基于人员基本信息，获取该人员在一段时间内参与的所有项目；
>> 4. 基于参与的所有项目，去重后获取项目的信息，包括：项目ID；项目名称；
>> 5. 基于人员和项目信息获取一段时间内该人员提交(commit)信息：提交ID；统计提交次数；
>> 6. 通过提交ID获取该提交涉及到的修改提交行数；
>> 7. 生成EXCEL表格进行分析；
>> 8. 通过与Tapd（项目管理），交互，获取每个人的缺陷数据；
>> 9. 获取tapd所有人的信息；
>> 10. 获取tapd所有项目信息。
> + ## 注意事项：
>> 使用的说协程开发，会涉及到相关锁资源的应用以及gitlab可以承受的并发数，如若修改，请联系最近的修改者。


> + ## 版本修改记录
>**版本号**  | **版本修改内容**  | **修改人**
>---|---|---
>v1.0.0 | 创建项目，实现gitlab的api调用，完成基本功能 | MassAdobe
>v1.0.1 | 生成excel并输出 | MassAdobe
>v1.0.2 | 增加Commit的次数 | MassAdobe
>v1.0.3 | 修改Sheet页，①KPI-VIEW；②KPI-SORT，修改KPI-SORT输出格式为每一行，方便排序 | MassAdobe
>v1.0.4 | 抽象封装了http的get请求，并运用到了所有方法中；抽出const字符；增加了gitignore | MassAdobe
>v1.0.5 | 增加与tapd交互的数据抽取 | MassAdobe
