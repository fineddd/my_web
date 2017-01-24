/**
 * Created by syz on 1/22/17.
 */
var UserTableEditable = function () {
    var nEditing = null;
    var oTable = null;
    //var nEditRow = null;
    return {
        //main function to initiate the module
        init: function () {
            // load data from server
            oTable = $('#user_table').dataTable({
                "paging": true,
                "lengthChange": true,
                "searching": false,
                "ordering": true,
                "info": true,
                "autoWidth": true,
                "aoColumnDefs": [{
                    'bSortable': false,
                    'aTargets': [6, 7]
                }]
            });

            $.post("/user/list",
                {
                },
                function(data, status) {
                    nEditing = null;
                    var ret = json_parse(data);
                    if (ret.errorCode == 0) {
                        //var aiNew = oTable.fnAddData(ret.users, true);
                        for (var i = 0; i < ret.users.length; i++) {
                            var lastLoginTime = '未登录过';
                            var status = 'offline';
                            if (ret.users[i].lastlogintime != 0)
                            {
                                lastLoginTime = new Date(ret.users[i].lastlogintime * 1000).toLocaleString().substr(0,21);
                            }
                            if (ret.users[i].status == 0) {
                                status = "online";
                            }
                            var strRight = '管理员';
                            if (ret.users[i].right == 0) {
                                strRight = "超级管理员";
                            }
                            var aiNew = oTable.fnAddData([ret.users[i].id, ret.users[i].name, strRight, ret.users[i].pfname,
                                lastLoginTime, status, '<a class="edit" href="">修改</a>', '<a class="delete" href="">删除</a>']);
                        }
                    } else {
                        alert(ret.str);
                        if (ret.errorCode == 2) {
                            window.location = "/";
                        }
                    }
                });

            $('#add_account').click(function (e) {
                e.preventDefault();
                if (nEditing !== null) {
                    alert("请等待其他操作完成！");
                    return;
                }
                $('#title_user_info').text("添加新用户");
                $('#btn_cancel').text("取消");
                $('#btn_commit').text("添加");
                $('#btn_commit').attr("user-mode", "new");
                $('#form_ck_pwd').hide();
                $('#form_pwd').show();
                $('#form_repwd').show();
                $('#in_username').val("");
                $('#in_status').hide();
                $('#in_pwd').val("");
                $('#in_repwd').val("");
                $('#user_info').modal('show');
/*
                $.post("/user/right",
                    {
                    },
                    function(data, status){
                        var ret = json_parse(data);
                        if (ret.errorCode == 0) {
                            $('#st_right').empty();
                            $('<option value="">请选择</option>').appendTo('#st_right');
                            for (var i = 0; i < ret.rights.length; i++) {
                                $('<option value="' + ret.rights[i].id + '">' + ret.rights[i].name + '</option>').appendTo('#st_right');
                            }
                            $('#st_right').attr("value", "");
                        } else {
                            alert(ret.str);
                            if (ret.errorCode == 2) {
                                if (ret.errorCode == 2) {
                                    window.location = "/";
                                }
                            }
                        }
                    });

                $.post("/platform/list",
                    {
                    },
                    function(data, status){
                        var ret = json_parse(data);
                        if (ret.errorCode == 0) {
                            var option = '';
                            for (var i = 0; i < ret.platforms.length; i++) {
                                option = option + '<option value="' + ret.platforms[i].id + '">' + ret.platforms[i].name + "</option>";
                            }
                            $("#platform_list").html(option);
                        } else {
                            alert(ret.str);
                            if (ret.errorCode == 2) {
                                window.location = "/";
                            } else {
                                return;
                            }
                        }
                    });*/
            });
        }
    };
}();