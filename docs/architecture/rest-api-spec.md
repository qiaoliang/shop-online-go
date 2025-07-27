# REST API Spec

```yaml
openapi: 3.0.0
info:
    title: 用户收货地址管理 API
    version: 1.0.0
    description: 提供用户收货地址的增删改查功能。

servers:
    - url: http://localhost:9090/v1
      description: 本地开发服务器

paths:
    /user/shipping-address/add:
        post:
            summary: 添加用户收货地址
            operationId: addShippingAddress
            requestBody:
                required: true
                content:
                    application/json:
                        schema:
                            $ref: "#/components/schemas/AddShippingAddressRequest"
            responses:
                "200":
                    description: 成功添加收货地址
                    content:
                        application/json:
                            schema:
                                $ref: "#/components/schemas/ApiResponse"
                "400":
                    description: 无效请求参数
                    content:
                        application/json:
                            schema:
                                $ref: "#/components/schemas/ApiResponse"

    /user/shipping-address/list:
        get:
            summary: 获取用户收货地址列表
            operationId: getShippingAddressList
            security:
                - bearerAuth: []
            responses:
                "200":
                    description: 成功获取收货地址列表
                    content:
                        application/json:
                            schema:
                                type: object
                                properties:
                                    code:
                                        type: integer
                                        example: 0
                                    msg:
                                        type: string
                                        example: "OK"
                                    data:
                                        type: array
                                        items:
                                            $ref: "#/components/schemas/Address"
                "401":
                    description: 未授权
                    content:
                        application/json:
                            schema:
                                $ref: "#/components/schemas/ApiResponse"

    /user/shipping-address/default:
        get:
            summary: 获取用户默认收货地址
            operationId: getDefaultShippingAddress
            security:
                - bearerAuth: []
            responses:
                "200":
                    description: 成功获取默认收货地址
                    content:
                        application/json:
                            schema:
                                type: object
                                properties:
                                    code:
                                        type: integer
                                        example: 0
                                    msg:
                                        type: string
                                        example: "OK"
                                    data:
                                        $ref: "#/components/schemas/Address"
                "401":
                    description: 未授权
                    content:
                        application/json:
                            schema:
                                $ref: "#/components/schemas/ApiResponse"

    /user/shipping-address/modify:
        post:
            summary: 修改用户收货地址
            operationId: modifyShippingAddress
            requestBody:
                required: true
                content:
                    application/json:
                        schema:
                            $ref: "#/components/schemas/ModifyShippingAddressRequest"
            responses:
                "200":
                    description: 成功修改收货地址
                    content:
                        application/json:
                            schema:
                                $ref: "#/components/schemas/ApiResponse"
                "400":
                    description: 无效请求参数
                    content:
                        application/json:
                            schema:
                                $ref: "#/components/schemas/ApiResponse"

    /user/shipping-address/delete:
        delete:
            summary: 删除用户收货地址
            operationId: deleteShippingAddress
            requestBody:
                required: true
                content:
                    application/json:
                        schema:
                            $ref: "#/components/schemas/DeleteShippingAddressRequest"
            responses:
                "200":
                    description: 成功删除收货地址
                    content:
                        application/json:
                            schema:
                                $ref: "#/components/schemas/ApiResponse"
                "400":
                    description: 无效请求参数
                    content:
                        application/json:
                            schema:
                                $ref: "#/components/schemas/ApiResponse"

    /user/shipping-address/set-default:
        post:
            summary: 设置用户默认收货地址
            operationId: setDefaultShippingAddress
            requestBody:
                required: true
                content:
                    application/json:
                        schema:
                            $ref: "#/components/schemas/SetDefaultShippingAddressRequest"
            responses:
                "200":
                    description: 成功设置默认收货地址
                    content:
                        application/json:
                            schema:
                                $ref: "#/components/schemas/ApiResponse"
                "400":
                    description: 无效请求参数
                    content:
                        application/json:
                            schema:
                                $ref: "#/components/schemas/ApiResponse"

components:
    securitySchemes:
        bearerAuth:
            type: http
            scheme: bearer
            bearerFormat: JWT

    schemas:
        Address:
            type: object
            properties:
                id:
                    type: string
                    description: 地址唯一标识
                userId:
                    type: string
                    description: 关联的用户 ID
                linkMan:
                    type: string
                    description: 联系人姓名
                mobile:
                    type: string
                    description: 联系手机号
                isDefault:
                    type: boolean
                    description: 是否为默认地址
                provinceStr:
                    type: string
                    description: 省份
                cityStr:
                    type: string
                    description: 城市
                areaStr:
                    type: string
                    description: 区域
                detailAddress:
                    type: string
                    description: 详细地址

        AddShippingAddressRequest:
            type: object
            required:
                - linkMan
                - mobile
                - provinceStr
                - cityStr
                - areaStr
                - detailAddress
            properties:
                linkMan:
                    type: string
                mobile:
                    type: string
                provinceStr:
                    type: string
                cityStr:
                    type: string
                areaStr:
                    type: string
                detailAddress:
                    type: string
                isDefault:
                    type: boolean
                    default: false

        ModifyShippingAddressRequest:
            type: object
            required:
                - id
            properties:
                id:
                    type: string
                linkMan:
                    type: string
                mobile:
                    type: string
                provinceStr:
                    type: string
                cityStr:
                    type: string
                areaStr:
                    type: string
                detailAddress:
                    type: string
                isDefault:
                    type: boolean

        DeleteShippingAddressRequest:
            type: object
            required:
                - id
            properties:
                id:
                    type: string

        SetDefaultShippingAddressRequest:
            type: object
            required:
                - id
            properties:
                id:
                    type: string

        ApiResponse:
            type: object
            properties:
                code:
                    type: integer
                    description: 响应代码 (0 表示成功)
                msg:
                    type: string
                    description: 响应消息
                data:
                    type: object
                    description: 响应数据 (可选)
```
