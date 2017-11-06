<div class="row subsystem-header">
    <div class="pull-left">
        <span style="font-size: 16px;">系统配置管理帮助信息</span>
    </div>
</div>
<div class="row subsystem-toolbar">
    <div class="col-sm-12 col-md-5 col-lg-3">
        <div class="pull-left">
            <span>导航信息</span>
        </div>
    </div>
    <div class="col-sm-12 col-md-7 col-lg-9" style="padding-left: 0px;">
        <div class="pull-left">
            <span>帮助信息</span>
        </div>
    </div>
</div>
<div class="row">
    <div class="col-sm-12 col-md-5 col-md-3">
        <div id="h-org-tree-info" style="border: #598f56 solid 1px;">
            <div id="h-help-list" class="col-sm-12 col-md-12 col-lg-12"
                 style="padding:8px 5px;overflow: auto">
                 <ul style="padding-left: 15px;">
                     <li onclick="javascript:document.getElementById('h-help-1').scrollIntoView()">1.系统概述</li>
                     <li onclick="javascript:document.getElementById('h-help-2').scrollIntoView()">2 功能模块设计</li>
                     <li onclick="javascript:document.getElementById('h-help-2.1').scrollIntoView()">2.1系统资源配置管理</li>
                     <li onclick="javascript:document.getElementById('h-help-2.1.1').scrollIntoView()">2.1.1 域定义管理详细介绍</li>
                     <li onclick="javascript:document.getElementById('h-help-2.1.1.1').scrollIntoView()">2.1.1.1 功能列表</li>
                     <li onclick="javascript:document.getElementById('h-help-2.1.1.2').scrollIntoView()">2.1.1.2 功能介绍</li>
                     <li onclick="javascript:document.getElementById('h-help-2.1.2').scrollIntoView()">2.1.2 组织架构详细介绍</li>
                     <li onclick="javascript:document.getElementById('h-help-2.1.2.1').scrollIntoView()">2.1.2.1 功能列表</li>
                     <li onclick="javascript:document.getElementById('h-help-2.1.2.1').scrollIntoView()">2.1.2.2 功能介绍</li>
                     <li onclick="javascript:document.getElementById('h-help-2.1.3').scrollIntoView()">2.1.3 菜单资源详细介绍</li>
                     <li onclick="javascript:document.getElementById('h-help-2.1.3.1').scrollIntoView()">2.1.3.1 功能列表</li>
                     <li onclick="javascript:document.getElementById('h-help-2.1.3.2').scrollIntoView()">2.1.3.2 功能介绍</li>
                     <li onclick="javascript:document.getElementById('h-help-2.2').scrollIntoView()">2.2 用户与权限</li>
                     <li onclick="javascript:document.getElementById('h-help-2.2.1').scrollIntoView()">2.2.1 用户管理</li>
                     <li onclick="javascript:document.getElementById('h-help-2.2.1.1').scrollIntoView()">2.2.1.1 功能列表</li>
                     <li onclick="javascript:document.getElementById('h-help-2.2.1.2').scrollIntoView()">2.2.1.2 功能列表</li>
                     <li onclick="javascript:document.getElementById('h-help-2.2.2').scrollIntoView()">2.2.2 角色管理</li>
                     <li onclick="javascript:document.getElementById('h-help-2.2.2.1').scrollIntoView()">2.2.2.1 功能列表</li>
                     <li onclick="javascript:document.getElementById('h-help-2.2.2.2').scrollIntoView()">2.2.2.2 功能介绍</li>
                     <li onclick="javascript:document.getElementById('h-help-2.2.3').scrollIntoView()">2.2.3 授权管理</li>
                     <li onclick="javascript:document.getElementById('h-help-2.2.3.1').scrollIntoView()">2.2.3.1 功能列表</li>
                     <li onclick="javascript:document.getElementById('h-help-2.2.3.2').scrollIntoView()">2.2.3.2 功能介绍</li>
                     <li onclick="javascript:document.getElementById('h-help-2.3').scrollIntoView()">2.3 系统审计</li>
                     <li onclick="javascript:document.getElementById('h-help-2.3.1').scrollIntoView()">2.3.1 日志查询</li>
                     <li onclick="javascript:document.getElementById('h-help-2.3.1.1').scrollIntoView()">2.3.1.1 功能列表</li>
                     <li onclick="javascript:document.getElementById('h-help-2.3.1.2').scrollIntoView()">2.3.1.2 功能介绍</li>
                 </ul>
            </div>
        </div>
    </div>
    <div class="col-sm-12 col-md-7 col-lg-9" style="padding-left: 0px;">
        <div id="h-help-details" style="border: #598f56 solid 1px;overflow: auto;padding: 15px">

            <strong id="h-help-1" style="height: 30px; line-height: 30px;">1.系统概述</strong><br/>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;集成开发平台，主要实现了用户与权限管理。在用户管理过程中，通过划分域，组织架构，来分层次的管理用户、权限信息。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;在权限控制方面，将每一个系统服务API，对应到系统菜单资源上，通过对菜单资源的授权管理，实现不同用户权限控制。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;系统组成部分：<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;第一部分：系统资源配置管理<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;第二部分：用户与权限配置管理<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;第三部分：系统操作审计管理<br/>
            </p>


            <strong id="h-help-2" style="height: 30px; line-height: 30px;">2 功能模块设计</strong><br/>
            <p style="line-height: 30px;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;系统管理主页信息如下图所示：<br/></p>

            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-10 col-md-10 col-lg-10" src="/static/images/help/image001.png"/>
            </div>

            <p style="line-height: 30px;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;上图是系统权限管理的主要功能结构图，系统提供了4个全局功能按钮。<br/></p>
            <p style="line-height: 30px;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;底部工具栏左起第一个按钮，用户返回到系统首页<br/></p>
            <p style="line-height: 30px;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;底部工具栏左起第二个按钮，是登陆用户的详细信息<br/></p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;底部工具栏左起第三个按钮，是安全退出按钮<br/></p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;底部工具栏左起第四个按钮，是返回子系统按钮，这个按钮将会从某一个具体的页面切换到子系统中，如从域定义页面中，切换到系统配置管理这个子系统。<br/></p>
            <strong id="h-help-2.1" style="height: 30px; line-height: 30px;">2.1系统资源配置管理</strong> <br/>

            <p style="line-height: 30px;">&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;系统资源配置管理部分由以下几个子页面系统组成：<br/></p>

            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	域定义管理<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	组织架构定义管理<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	菜单资源的配置管理<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;三个子页面系统的关系，如下图所示：<br/>
            </p>

            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image003.png"/>
            </div>

            <strong id="h-help-2.1.1" style="height: 30px; line-height: 30px;">2.1.1 域定义管理详细介绍</strong><br/>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;主要是在系统中开辟了一个独立的空间，在这个空间中，可以构建自己的组织架构，用户，角色。 此外，可以通过配置域的共享，来实现域与域之间的信息共享。用户需要被授予查询域信息权限，才能查看到系统中域列表信息.<br/>
            </p>
            <strong id="h-help-2.1.1.1" style="height: 30px; line-height: 30px;">
                2.1.1.1 功能列表<br/>
            </strong>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	新增<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	编辑<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	删除<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	共享管理<br/>
            </p>

            <strong id="h-help-2.1.1.2" style="height: 30px; line-height: 30px;">2.1.1.2 功能介绍</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;新增：用于新增域信息，在新增域时，需要提供3个字段信息。用户需要被授予“新增域信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image004.png"/>
            </div>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;域编码：用于在系统中标识域的一个标记，整个系统中，域编码必须具有唯一性。域编码长度必须由1-30位字母、数字组成<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;域描述：用户更清晰的描述域的信息。描述信息必须由1-30位汉字，字母，数字组成。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;状态：状态分为2种，一种是正常，一种是锁定。当域被锁定后，域中的资源无法被访问。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;编辑：用于编辑已经存在的域信息。用户需要被授予“编辑域信息按钮”权限，才能进行此项操作。<br/>
            </p>

            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image006.png"/>
            </div>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;选中域信息列表中的某一个域（同一时点，只能编辑一个域），域编码无法被编辑，只能修改域描述和域状态。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;删除：用于删除已经存在的域信息，系统会首先检测用户是否能够对这个域进行读写，如果没有写权限，则禁止用户删除这个域。用户需要被授予“删除域信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;1. 用户无法删除自己所在的域<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;2. 用户无法删除没有被授予读写权限的域<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;在点击删除按钮时，首先会弹出提示确认删除窗口，一旦点击确认，系统便会检查用户权限，检查依赖关系。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image008.png"/>
            </div>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;共享管理：共享管理操作，可以将这个域的读写权限授予给另一个域，这个按钮是一个链接。用户需要被授予“共享域管理”权限，才能进行此项操作。<br/>
            </p>

            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-10 col-md-10 col-lg-10" src="/static/images/help/image010.png"/>
            </div>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;上图是域共享配置界面。左边栏显示的是被共享的域详细信息。右边栏，用于配置接收共享的目标域。如果用户对被共享域有读写权限，则可以操作右侧的新增，编辑，删除功能，如果只有读权限，则不能操作右侧的3个功能按钮。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;右侧的所属域，可以切换被共享域，这是一个下拉框，只能显示用户可以访问到的域列表信息，没有被授权共享的域，不会在这个下拉框中出现。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;用户需要被授予“查询共享域信息”权限，才能看到被共享授权信息列表。但是有一个特例，admin用户不受共享域的约束，admin用户可以访问任何域空间。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;这个页面中，提供了如下几个功能点：<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	新增<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	编辑<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	删除<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;新增：就是将左侧显示的被共享域授权共享给新的目标域。用户需要有“新增共享域信息按钮”权限，才能进行此项操作。<br/>
            </p>

            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image012.png"/>
            </div>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;共享目标：只会显示剩余没有被授权共享的域。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;共享模式：分为两种，一种是只读，另一种是读写。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;编辑：可以对域的共享模式进行编辑，如果用户能够对被共享的域进行读写，则可以修改共享目标的授权模式。用户需要有“更新共享域信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image014.png"/>
            </div>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;用户选择右侧列表框中的共享目标，对其进行编辑。在同一时点，只能对一个共享目标进行编辑。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;删除：如果用户能够对被共享的域进行读写，则可以删除被共享域的共享目标。用户需要有“删除共享域信息按钮”权限，才能进行此项操作。<br/>
            </p>

            <strong id="h-help-2.1.2" style="height: 30px; line-height: 30px;">2.1.2 组织架构详细介绍</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;每一个域中，都有自己独立的组织架构。也即是说，组织架构存在于特定的域中，用于构建企业，团队内部的一种关系。是用户的载体，所有的用户，都必须指定明确的归属组织。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;用户需要被授予“查询组织架构信息”权限，才能查询组织架构详细信息。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-10 col-md-10 col-lg-10" src="/static/images/help/image016.png"/>
            </div>

            <strong id="h-help-2.1.2.1" style="height: 30px; line-height: 30px;">2.1.2.1 功能列表</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	新增<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	编辑<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	删除<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	导入<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	导出<br/>
            </p>

            <strong id="h-help-2.1.2.2" style="height: 30px; line-height: 30px;">2.1.2.2 功能简介</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;新增：用于添加新的组织进入组织架构体系中。用户需要被授予“新增组织架构信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image018.png"/>
            </div>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;新增操作，需要录入4个字段信息，<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;第一个是组织部门代码，编码必须是1-30位字母或数字组成<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;第二个是组织部门名称，名称必须是1-30个汉字，字母，数字组成<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;第三个是这个组织代码所属域，<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;第四个是上级组织部门代码。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;上述4个字段都不能为空。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;删除：用于删除组织架构中的某个节点以及该节点下的所有组织信息。在右侧的机构详细信息中选择要删除的机构，会弹出提示框，点击确认后，将会删除这个组织，以及这个组织下的所有组织信息，如果这个组织下边已经配置有用户，则将会取消本次删除操作。用户需要被授予“删除组织架构信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;编辑：用户对组织架构中的某个组织进行编辑操作。在编辑组织信息过程中，这个组织的上级组织不能变更成这个组织的下级组织。也就是父子关系不能对调。用户需要被授予“更新组织架构信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image020.png"/>
            </div>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;编辑框中的组织机构代码不能被修改。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;导出：将会把指定域中的所有机构信息导出到excel中。用户需要被授予“导出组织架构信息按钮”权限，才能进行此项操作<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;导入：将会根据模板填充信息，将组织架构导入到系统中。用户需要被授予“导出组织架构信息按钮”权限，才能进行此项操作。模板信息必须按照规定填写，在导入的过程中，有两种选择：<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	增量导入，覆盖重复。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;这种模式中，将会根据模板中的数据来构建组织架构，在构建的过程中，如果模板中出现了存量的组织信息，将会覆盖原有系统中信息，使用模板中的组织信息替换。判断增量，存量的依据是主键ID，也就是机构编码。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	全量导入，删除历史<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;这种模式会首先删除已有的数据，然后根据模板中数据，填充组织架构信息。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	增量导入，丢弃重复<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;这种模式中，不会对已有的数据做任何操作只会将新增的信息导入进去，根据主键ID，也就是机构号，来判断是否是增量数据。也即是说，如果主键不变，其他唯独变化，也会被认定成存量，而非增量，对于模板中出现的存量信息，将会被丢弃。<br/>
            </p>

            <strong id="h-help-2.1.3" style="height: 30px; line-height: 30px;">2.1.3 菜单资源详细介绍</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;菜单资源是系统提供的功能点，被所有的域共享。这块资源比较特殊，他是系统的公共资源，也是系统的核心部件。这部分资源只针对超级管理员。当整个系统有新的功能加入，调整主题颜色，菜单名称时，才需要使用这部分功能。对于业务操作与管理上，无需授予这部分权限。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;用户需要被授予“查询资源信息“权限，才能查看到系统中的菜单资源信息。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-10 col-md-10 col-lg-10" src="/static/images/help/image022.png"/>
            </div>

            <strong id="h-help-2.1.3.1" style="height: 30px; line-height: 30px;">2.1.3.1 功能列表</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	新增<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	编辑<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	配置<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	删除<br/>
            </p>

            <strong id="h-help-2.1.3.2" style="height: 30px; line-height: 30px;">2.1.3.2 功能简介</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;新增：新增菜单资源，用户像系统资源列表中新增菜单资源信息，用户需要被授予“新增资源信息按钮“权限，才能进行此项操作。<br/>
            </p>

            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image024.png"/>
            </div>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;所属主题：默认选择青春风格，系统目前不支持分割切换，主要原因是系统中目前只配置了一套风格，即青春风格。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;菜单类别：<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	第一种是首页系统，用于构建一个单独的业务系统的操作，比如新增调度模块，报表模块，则需要先建立一个首页系统菜单。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	第二种是子页系统，在指定的子模块系统中创建一个子页面系统，比如在报表系统模块中，创建一个过滤器管理页面。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	第三种是功能按钮。按钮资源不会直接显示在主页或子页面中，而是作为一种服务，嵌套在具体的页面中，通过配置按钮资源的权限，可以方便的控制用户对于配置表的增删查改的权限。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	第四种是虚拟节点，这类主要用于创建一个虚拟的节点，用于对下边子资源的一个打包处理。这类资源没有实际的操作特性，仅仅只是对子资源的一个包装，汇集，让资源树结构更清晰、明了。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;资源编码：用于标识这个资源的唯一ID号；<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;资源名称：用于对这个资源的具体描述；<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;路由信息：填写API服务地址。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;样式属性：系统提供了3中样式，分别是小正方形，长方形，大正方向三种。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;图标路径：是资源的图标地址<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;图标色彩：是资源的色彩值<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;分组编号：是资源所在的列，通常从左至右分为三列。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;排序号：是资源在指定的列中顺序。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;备注：对于菜单类型为非首页系统时，还需要填写上级资源信息。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;编辑：用于对菜单资源的名称进行调整。用户需要被授予“编辑资源信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image026.png"/>
            </div>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;在编辑过程中，只能对资源名称进行编辑。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;删除：用户删除菜单资源信息，当点击删除时，将会连同这个资源以及其子资源一起删除。如果被删除的资源是系统资源，则会提示删除失败，系统资源禁止删除。用户需要被授予“编辑资源信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;配置：用户修改资源的主题配置信息，用户需要被授予“配置主题信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image028.png"/>
            </div>

            <strong id="h-help-2.2" style="height: 30px; line-height: 30px;">2.2 用户与权限</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;用户与权限管理部分，主要包含三部分功能，分别是：<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	用户管理<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	角色管理<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	授权管理<br/>
            </p>

            <strong id="h-help-2.2.1" style="height: 30px; line-height: 30px;">2.2.1 用户管理</strong>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;用户管理子页面功能，主要完成了对系统中用户信息的配置管理。用户需要被授予“查询用户信息”权限，才能进行此项操作。<br/>
            </p>

            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-10 col-md-10 col-lg-10" src="/static/images/help/image030.png"/>
            </div>

            <strong id="h-help-2.2.1.1" style="height: 30px; line-height: 30px;">2.2.1.1 功能列表</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	新增<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	编辑<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	删除<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	改密<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	解锁<br/>
            </p>

            <strong id="h-help-2.2.1.2" style="height: 30px; line-height: 30px;">2.2.1.2 功能简介</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;新增：用于新增用户信息，用户需要被授予“新增用户信息按钮”权限，才能进行此项操作。<br/>
            </p>

            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image032.png"/>
            </div>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;账号：用户登陆系统的唯一代号。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;名称：用户描述账户的具体信息。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;密码：用户使用账户登陆系统所需要的密码。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;确认密码：确认登陆密码。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;邮箱：用户使用接受信息的邮箱。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;手机号：用户联系方式。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;所属域：用户归属的域<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;机构：用户归属的机构<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;编辑：用于编辑已有的用户信息，用户需要被授予“编辑用户信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image034.png"/>
            </div>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;编辑操作，提供了对用户名称，邮箱，手机号，机构的修改。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;删除：用于删除系统中已有的用户信息，用户只能删除自己能够进行读写的域中的用户，但是不能删除用户自己和admin用户。用户需要被授予“删除用户信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;改密：用户修改指定用户的密码，改密操作时，根据指定的要求，填写符合格式的信息，即可修改密码。用户需要被授予“修改用户密码按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image036.png"/>
            </div>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;解锁：用于管理用户的状态信息，只需要选择用户状态即可。用户需要被授予“修改用户状态按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image038.png"/>
            </div>

            <strong id="h-help-2.2.2" style="height: 30px; line-height: 30px;">2.2.2角色管理</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;角色在系统权限控制中，是拥有某些菜单资源的个体。每一个角色，都拥有自己能够访问到的资源。用户需要被授予“角色查询信息”权限，才能查看角色列表信息。<br/>
            </p>

            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-10 col-md-10 col-lg-10" src="/static/images/help/image040.png"/>
            </div>

            <strong id="h-help-2.2.2.1" style="height: 30px; line-height: 30px;">2.2.2.1 功能列表</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	新增<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	编辑<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	删除<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	资源管理<br/>
            </p>

            <strong id="h-help-2.2.2.2" style="height: 30px; line-height: 30px;">2.2.2.2 功能简介</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;新增：用户新增角色编码信息，用户需要被授予“新增角色信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image042.png"/>
            </div>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;新增角色信息接收4个字段，<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;角色编码：在所在的域中具有唯一性。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;角色名称：用户描述具体的描述角色编码<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;状态：分为两种情况：正常，失效。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;所属域：角色归属域。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;编辑：修改已经存在角色的名称，状态信息，用户需要被授予“更新角色信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image044.png"/>
            </div>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;删除：用户删除已经存在的角色信息，用户需要被授予“删除角色信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;资源管理：配置角色拥有的菜单资源信息。这是一个跳转连接，用户如果拥有“角色资源管理”权限，当点击这个按钮后，会跳转到一个菜单资源配置页面。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-10 col-md-10 col-lg-10" src="/static/images/help/image046.png"/>
            </div>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;左起第一列，是角色的详细信息。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;左起第二列，是这个角色已经拥有的菜单资源信息。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;左起第三列，是这个角色尚未拥有的菜单资源信息。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;通过选中菜单资源，点击撤销，则会将资源以及这个资源的子资源从角色中撤销。当点击授权时，将会将这个资源以及这个资源的子资源授予这个角色。<br/>
            </p>

            <strong id="h-help-2.2.3" style="height: 30px;line-height: 30px;">2.2.3 授权</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;授权子页面，给用户赋予相应的角色，使这个用户有权访问系统中角色拥有的资源。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-10 col-md-10 col-lg-10" src="/static/images/help/image048.png"/>
            </div>

            <strong id="h-help-2.2.3.1" style="height: 30px;line-height: 30px;">2.2.3.1 功能列表</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	授权<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	删除<br/>
            </p>

            <strong id="h-help-2.2.3.2" style="height: 30px;line-height: 30px;">2.2.3.2 功能简介</strong><br/>
            <p style="line-height: 30px">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;授权：在用户信息列表中选择一个用户，点击授权按钮，将会弹出选择框，这个框中会显示这个用户尚未拥有的角色。用户需要被授予“授予权限按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image050.png"/>
            </div>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;选择需要授予给这个用户的角色，点击提交，即完成了授权操作。<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;删除：在右边的用户已经拥有角色栏中，点击删除，将会删除这个用户已经拥有的角色信息。<br/>
            </p>

            <strong id="h-help-2.3" style="height: 30px;line-height: 30px;">2.3 系统审计</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;系统设计部分，主要实现了对操作记录的查询。由日志查询模块组成。<br/>
            </p>

            <strong id="h-help-2.3.1" style="height: 30px;line-height: 30px;">2.3.1 日志查询</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;通过日志查询功能，对登陆系统用户的操作进行分析，解决一些恶意的破坏问题，为后期追责提供依据。用户需要被授予“查询操作日志”权限，才能进行查看到操作记录。用户只能查到到自己域中的操作记录。<br/>
            </p>

            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-10 col-md-10 col-lg-10" src="/static/images/help/image052.png"/>
            </div>


            <strong id="h-help-2.3.1.1" style="height: 30px;line-height: 30px;">2.3.1.1 功能列表</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	搜索<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	下载<br/>
            </p>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;	查询请求参数<br/>
            </p>

            <strong id="h-help-2.3.1.2" style="height: 30px;line-height: 30px;">2.3.1.2 功能简介</strong><br/>

            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;搜索，可以根据账号，操作日期，进行操作记录查询，用户需要被授予“搜索日志信息按钮”权限，才能进行此项操作。<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image054.png"/>
            </div>
            <p style="line-height: 30px;">
            &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;下载：用户下载操作记录到excel中。用户需要被授予“下载操作日志按钮”权限，才能此项操作。<br/>
            </p>
            <p style="line-height: 30px;">
                &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;查询请求参数：双击请求发送列中的数据，会弹出显示框，显示用户发送到服务器的请求参数：<br/>
            </p>
            <div class="col-sm-12 col-md-12 col-lg-12">
                <img class="col-sm-6 col-md-6 col-lg-6" src="/static/images/help/image056.png"/>
            </div>
        </div>
    </div>
</div>

<script type="text/javascript">

    $(document).ready(function(){
        var hwindow = document.documentElement.clientHeight;
        $("#h-org-tree-info").height(hwindow-130);
        $("#h-help-details").height(hwindow-160);
        $("#h-help-list").height(hwindow-160);


        $("#h-help-list").find("li").bind("mouseover",function () {
            $(this).css({
                "cursor":"pointer",
                "color":"blue",
            })
        }).bind("mouseout",function () {
            $(this).css({
                "cursor":"",
                "color":"black",
            });
        }).css({
            "height":"30px",
            "line-height":"30px",
        })
    });
</script>