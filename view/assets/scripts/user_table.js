/**
 * Created by syz on 1/22/17.
 */
var UserTableEditable = function () {
    var nEditing = null;
    var oTable = null;
    var nEditRow = null;

    $('#user_table').on('click', 'a.delete', function(e) {
        e.preventDefault();
        if (nEditing !== null) {
            alert("请等待其他操作完成 !!!");
            return;
        }
        if (confirm("是否删除？") == false) {
            return;
        }
        var nRow = $(this).parents('tr')[0];
        var aData = oTable.fnGetData(nRow);
        nEditing = nRow;
        $.post("/user/del",
            {
                id : aData[0],
                pfid : aData[8]
            },
            function(data, status){
                var ret = json_parse(data);
                if (ret.errorCode == 0) {
                    oTable.fnDeleteRow(nRow);
                } else {
                    alert(ret.str);
                    if (ret.errorCode == 2) {
                        window.location = "/";
                    }
                }
                nEditing = null;
            });
    });

    $('#user_table').on('click', 'a.edit', function(e) {
        e.preventDefault();
        if (nEditing !== null) {
            alert("请等待其他操作完成 !!!");
            return;
        }

        var nRow = $(this).parents('tr')[0];
        var aData = oTable.fnGetData(nRow);
        nEditRow = nRow;
        $('#title_user_info').text("修改用户信息");
        $('#btn_cancel').text("取消");
        $('#btn_commit').text("修改");
        $('#btn_commit').attr("user-mode", "old");
        $('#btn_commit').attr("user-id", aData[0]);
        $('#form_pwd').hide();
        $('#form_repwd').hide();
        $('#form_ck_pwd').show();
        $('#in_status').hide();

        $('#in_username').val(aData[1]);

        $('#btn_ck_pwd').prop('checked', false);
        $('#in_pwd').val("");
        $('#in_repwd').val("");

        $('#user_info').modal('show');

        $.post("/right/list",
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
                    $('#st_right').attr("value", aData[9]);
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
                    var option;
                    for (var i = 0; i < ret.platforms.length; i++) {
                        option = option + '<option value="' + ret.platforms[i].id + '">';
                        option = option + ret.platforms[i].name + "</option>";
                    }
                    $("#platform_list").html(option);
                    $("#platform_list").val(aData[8]);
                } else {
                    alert(ret.str);
                    if (ret.errorCode == 2) {
                        window.location = "/";
                    } else {
                        return;
                    }
                }
            });
    });

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
                    'aTargets': [5, 6, 7]
                }],
                "aoColumns": [null, null, null, null, null, null, null, null, {"bVisible": false}, {"bVisible": false}]
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
                            if (ret.users[i].lastlogintime != 0)
                            {
                                lastLoginTime = new Date(ret.users[i].lastlogintime * 1000).toLocaleString().substr(0,21);
                            }

                            var aiNew = oTable.fnAddData([ret.users[i].id, ret.users[i].name, ret.users[i].rightname, ret.users[i].pfname,
                                lastLoginTime, ret.users[i].status, '<a class="edit" href="">修改</a>', '<a class="delete" href="">删除</a>', ret.users[i].pfid, ret.users[i].right]);
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

                $.post("/right/list",
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
                    });
            });

            $('#btn_commit').click(function (e) {
                e.preventDefault();

                if ($('#st_right').find(":selected").val() == "") {
                    alert("请选择权限组");
                    return false;
                }

                if ($(this).attr("user-mode") == "new") {
                    if ($('#in_pwd').val() != $('#in_repwd').val()) {
                        alert("两次密码不一致！");
                        return false;
                    }
                    if ($('#in_pwd').val().length < 6) {
                        alert("密码长度过短！");
                        return false;
                    }
                    nEditing = 0;
                    $.post("/user/add",
                        {
                            pfid : $('#platform_list').find(":selected").val(),
                            name : $('#in_username').val(),
                            password : $('#in_pwd').val(),
                            right : $('#st_right').find(":selected").val(),
                        },
                        function(data, status) {
                            var ret = json_parse(data);
                            if (ret.errorCode == 0) {
                                var aiNew = oTable.fnAddData([ret.id, $('#in_username').val(), $('#st_right').find(":selected").text(),
                                    $('#platform_list').find(":selected").text(), "未登录过", 0, '<a class="edit" href="">修改</a>', '<a class="delete" href="">删除</a>',
                                    $('#platform_list').find(":selected").val(), $('#st_right').find(":selected").val()]);
                            } else {
                                alert(ret.str);
                                if (ret.errorCode == 2) {
                                    window.location = "/";
                                }
                            }
                            nEditing = null;
                        });
                } else {
                    nEditing = 0;
                    if ($('#btn_ck_pwd').prop('checked')) {
                        if ($('#in_pwd').val() != $('#in_repwd').val()) {
                            alert("两次密码不一致！");
                            return false;
                        }
                        if ($('#in_pwd').val().length < 6) {
                            alert("密码长度过短！");
                            return false;
                        }
                    }
                    $.post("/user/update",
                        {
                            id : $(this).attr("user-id"),
                            pfid : $('#platform_list').find(":selected").val(),
                            name : $('#in_username').val(),
                            right : $('#st_right').find(":selected").val(),
                            changepassword : $('#btn_ck_pwd').prop('checked'),
                            password : $('#in_pwd').val()
                        },
                        function(data, status) {
                            var ret = json_parse(data);
                            if (ret.errorCode == 0) {
                                oTable.fnUpdate($('#platform_list').find(":selected").text(), nEditRow, 3, false);
                                oTable.fnUpdate($('#in_username').val(), nEditRow, 1, false);
                                oTable.fnUpdate($('#st_right').find(":selected").text(), nEditRow, 2, false);
                                oTable.fnUpdate($('#platform_list').find(":selected").val(), nEditRow, 8, false);
                                oTable.fnUpdate($('#st_right').find(":selected").val(), nEditRow, 9, false);
                                oTable.fnDraw();
                            } else {
                                alert(ret.str);
                                if (ret.errorCode == 2) {
                                    window.location = "/";
                                }
                            }
                            nEditing = null;
                        });
                }
            });
            $('#btn_ck_pwd').click(function (e) {
                if ($(this).prop('checked')) {
                    $('#form_pwd').show();
                    $('#form_repwd').show();
                } else {
                    $('#form_pwd').hide();
                    $('#form_repwd').hide();
                }
            });
        }
    };
}();