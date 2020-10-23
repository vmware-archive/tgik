json.array!(@widgets) do |widget|
  json.extract! widget, :id, :name, :description, :stock
  json.url widget_url(widget, format: :json)
end
