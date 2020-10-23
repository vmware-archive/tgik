class CreateWidgets < ActiveRecord::Migration[5.1]
  def change
    create_table :widgets do |t|
      t.string :name
      t.text :description
      t.integer :stock

      t.timestamps
    end
  end
end
