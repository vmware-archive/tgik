require 'test_helper'

class WidgetsControllerTest < ActionController::TestCase
  setup do
    @widget = widgets(:one)
  end

  test "should get index" do
    get :index
    assert_response :success
  end

  test "should get new" do
    get :new
    assert_response :success
  end

  test "should create widget" do
    assert_difference('Widget.count') do
      post :create,
        params: {
          widget: {
            description: @widget.description,
            name: @widget.name,
            stock: @widget.stock
          }
        }
    end

    assert_redirected_to widget_path(@widget.id + 1)
  end

  test "should show widget" do
    get :show, params: { id: @widget }
    assert_response :success
  end

  test "should get edit" do
    get :edit, params: { id: @widget }
    assert_response :success
  end

  test "should update widget" do
    patch :update,
      params: {
        id: @widget,
        widget: {
          description: @widget.description,
          name: @widget.name,
          stock: @widget.stock
        }
      }
    assert_redirected_to widget_path(@widget)
  end

  test "should destroy widget" do
    assert_difference('Widget.count', -1) do
      delete :destroy, params: { id: @widget }
    end

    assert_redirected_to widgets_path
  end
end
